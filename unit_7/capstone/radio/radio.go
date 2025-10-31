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
	time         time.Time
	callSign     string
	destination  string
	data         mars.SensorData
	retransmits  uint8
	broadcasters []string
}

func (m Message) hasBroadcast(callSign string) bool {
	for _, b := range m.broadcasters {
		if callSign == b {
			return true
		}
	}
	return false
}

func (m Message) Hash() string {
	return fmt.Sprintf("%v %v %v", m.callSign, m.time, m.data)
}

type Buffer struct {
	buffer []Message
}

var allRadiosMutex sync.RWMutex
var allRadios []*Radio = make([]*Radio, 0)

type Antenna struct {
	// DataRate is the number of bytes that can be sent per second
	DataRate       uint
	SignalRange    uint
	Distancer      Distancer
	broadcastMutex sync.Mutex
}

type Distancer interface {
	Distance(other Distancer) uint
}

type Radio struct {
	callSign         string
	receivingChannel chan Message
	inBufferMutex    sync.Mutex
	inBuffer         map[string]Message
	*Antenna
}

func (r *Radio) Inbox() []Message {
	r.inBufferMutex.Lock()
	defer r.inBufferMutex.Unlock()
	var buffer []Message = make([]Message, 0)
	for _, message := range r.inBuffer {
		buffer = append(buffer, message)
	}
	fmt.Printf("%v - inbox has %v messages in it\n", r.callSign, len(buffer))
	r.inBuffer = make(map[string]Message)
	return buffer
}

func (r *Radio) SendData(destination string, data mars.SensorData) {
	r.broadcastMessage(Message{
		time:         time.Now(),
		callSign:     r.callSign,
		broadcasters: []string{r.callSign},
		destination:  destination,
		data:         data,
		retransmits:  5,
	})
}

func (r *Radio) broadcastMessage(message Message) {
	log.Printf("%v transmitting", r.callSign)
	r.dataRateDelay(message)
	message.broadcasters = append(message.broadcasters, r.callSign)
	for _, other := range r.getRadiosInRange() {
		log.Printf("%v - hopefully %v will receive", r.callSign, other.callSign)
		go r.sendMessageWithDelays(message, other)
	}
}

func (r *Radio) dataRateDelay(message Message) {
	r.broadcastMutex.Lock()
	defer r.broadcastMutex.Unlock()
	size := uint(unsafe.Sizeof(message))
	ms := 1000 * float64(size) / float64(r.DataRate)
	log.Printf("%v - broadcast datarate is %v b/s so it'll take %v seconds to transmit a %v byte message", r.callSign, r.DataRate, ms/1000, size)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (r *Radio) lightSpeedDelay(receiver *Radio) {
	const lightSpeed = 100 * 299.792 // km/ms
	distance := r.Distancer.Distance(receiver.Distancer)
	ms := float64(distance) / lightSpeed
	log.Printf("%v - lightspeed delay will be %v milliseconds to get to %v", r.callSign, ms, receiver.callSign)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (r *Radio) sendMessageWithDelays(message Message, receiver *Radio) {
	r.lightSpeedDelay(receiver)
	log.Printf("%v - after long wait %v should have received a message", r.callSign, receiver.callSign)
	receiver.receivingChannel <- message
}

func (r *Radio) listen() {
	for message := range r.receivingChannel {
		log.Printf("%v - received message %v", r.callSign, message)
		if message.destination == r.callSign {
			// Sweet this message was meant for us! Let's add it to our incoming buffer
			r.inBufferMutex.Lock()
			r.inBuffer[message.Hash()] = message
			r.inBufferMutex.Unlock()
		} else if message.hasBroadcast(r.callSign) {
			continue
		} else if message.retransmits > 0 {
			message.retransmits -= 1
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
		inBuffer:         make(map[string]Message),
	}
	go radio.listen()
	allRadios = append(allRadios, radio)
	return radio
}
