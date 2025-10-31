package radio

import (
	"capstone/mars"
	"fmt"
	"log"
	"sync"
	"time"
	"unsafe"
)

/*
I think that I'd like to have a frequency space where the underlying implementation is channels but
the user doesn't know that.

We want to find live on Mars, so we'll send several rovers down to search for it, but we need to
know when life is found. In every cell in the grid, assign some likelihood of life, a random number
between 0 and 1000. If a rover finds a cell with a life value above 900, it may have found life and
it must send a radio message back to Earth.

Unfortunately, it's not always possible to send a message immediately because the relay satellite
is not always above the horizon. Implement a buffer goroutine that receives messages sent from the
rover and buffers them into a slice until they can be sent back to Earth.

Implement Earth as a goroutine that receives messages only occasionally (in reality for a couple of
hours every day, but you might want to make the interval a little shorter than that). Each message
should contain the coordinates of the cell where the life might have been found, and the life
value itself.

TODO Rovers need names
You may also want to give a name to each of your rovers and include that in the message so you can
see which rover sent it. It's also helpful to include the name in the log messages printed by the
rovers so you can track the progress of each one.

Set your rovers free to search and see what they come up with!
*/

/*
How do I model a radio?
*/

type Message struct {
	time        time.Time
	callSign    string
	destination string
	data        mars.SensorData
}

func (m Message) Hash() string {
	return fmt.Sprintf("%v %v %v", m.callSign, m.time, m.data)
}

func (m Message) String() string {
	return fmt.Sprintf(
		"Callsign: %v - Signs of life %v at (%v, %v)",
		m.callSign,
		m.data.LifeSigns,
		m.data.Location.X,
		m.data.Location.Y,
	)
}

type Buffer struct {
	mu     sync.Mutex
	buffer map[string]Message
}

func (b *Buffer) Add(message Message) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.buffer[message.Hash()] = message
}
func (b *Buffer) Pop() []Message {
	b.mu.Lock()
	defer b.mu.Unlock()
	contents := make([]Message, 0)
	for hash, message := range b.buffer {
		contents = append(contents, message)
		delete(b.buffer, hash)
	}
	return contents
}

func (b *Buffer) Contains(message Message) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	for hash := range b.buffer {
		if hash == message.Hash() {
			return true
		}
	}
	return false
}
func (b *Buffer) DumpOlderThan(t time.Time) {}

func NewBuffer() *Buffer {
	return &Buffer{buffer: make(map[string]Message)}
}

var allRadiosMutex sync.RWMutex
var allRadios []*Radio = make([]*Radio, 0)

type Antenna struct {
	// DataRate is the number of bytes that can be sent per second
	DataRate       uint
	SignalRange    float64
	Distancer      Distancer
	broadcastMutex sync.Mutex
}

type Distancer interface {
	Distance(other Distancer) float64
}

type Radio struct {
	callSign         string
	receivingChannel chan Message
	inBuffer         Buffer
	outBuffer        Buffer
	*Antenna
}

func (r *Radio) Inbox() []Message {
	return r.inBuffer.Pop()
}

func (r *Radio) SendData(destination string, data mars.SensorData) {
	r.broadcastMessage(Message{
		time:        time.Now(),
		callSign:    r.callSign,
		destination: destination,
		data:        data,
	})
}

func (r *Radio) broadcastMessage(message Message) {
	log.Printf("%v transmitting", r.callSign)
	r.outBuffer.Add(message)
	r.dataRateDelay(message)
	for _, other := range r.getRadiosInRange() {
		go r.sendMessageWithDelays(message, other)
	}
}

func (r *Radio) dataRateDelay(message Message) {
	r.broadcastMutex.Lock()
	defer r.broadcastMutex.Unlock()
	size := uint(unsafe.Sizeof(message))
	ms := 1000 * float64(size) / float64(r.DataRate)
	// log.Printf("%v - broadcast datarate is %v b/s so it'll take %v seconds to transmit a %v byte message", r.callSign, r.DataRate, ms/1000, size)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (r *Radio) lightSpeedDelay(receiver *Radio) {
	const lightSpeed = 100 * 299.792 // km/ms
	distance := r.Distancer.Distance(receiver.Distancer)
	ms := float64(distance) / lightSpeed
	// log.Printf("%v - lightspeed delay will be %v milliseconds to get to %v", r.callSign, ms, receiver.callSign)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (r *Radio) sendMessageWithDelays(message Message, receiver *Radio) {
	r.lightSpeedDelay(receiver)
	// log.Printf("%v - after long wait %v should have received a message", r.callSign, receiver.callSign)
	receiver.receivingChannel <- message
}

func (r *Radio) listen() {
	for message := range r.receivingChannel {
		log.Printf("%v - received message %v", r.callSign, message)
		if message.destination == r.callSign {
			r.inBuffer.Add(message)
		} else if r.outBuffer.Contains(message) {
			continue
		} else {
			go r.broadcastMessage(message)
		}
		// If this message was meant for someone else, how can we send it and not have it bounce endlessly?
	}
}

func (r *Radio) getRadiosInRange() []*Radio {
	allRadiosMutex.RLock()
	defer allRadiosMutex.RUnlock()
	radios := make([]*Radio, 0)
	for i := range allRadios {
		distance := r.Distancer.Distance(allRadios[i].Distancer)
		if r != allRadios[i] && distance <= r.SignalRange && distance <= allRadios[i].SignalRange {
			radios = append(radios, allRadios[i])
		}
	}
	return radios
}

type Receiver interface {
	Receive(message Message)
}

// NewRadio returns a new radio.
// It may be used concurrently by different goroutines.
func NewRadio(callSign string, antenna *Antenna) *Radio {
	allRadiosMutex.Lock()
	defer allRadiosMutex.Unlock()
	radio := &Radio{
		callSign:         callSign,
		Antenna:          antenna,
		receivingChannel: make(chan Message),
		inBuffer:         *NewBuffer(),
		outBuffer:        *NewBuffer(),
	}
	go radio.listen()
	allRadios = append(allRadios, radio)
	return radio
}
