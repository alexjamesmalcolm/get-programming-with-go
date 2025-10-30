package radio

import (
	"capstone/mars"
	"fmt"
	"log"
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
	callSign string
	data     mars.SensorData
}

type Buffer struct {
	buffer []Message
}

var allRadios []*Radio = make([]*Radio, 0)

type Antenna struct {
	SignalRange int
	Distancer
}

type Distancer interface {
	Distance(other Distancer) int
}

type Radio struct {
	callSign string
	receiver Receiver
	Antenna
}

func (r *Radio) SendMessage(data mars.SensorData) {
	log.Printf("%v transmitting", r.callSign)
	for _, other := range r.getRadiosInRange() {
		log.Printf("%v is in range of %v", other.callSign, r.callSign)
		other.receiver.Receive(Message{
			callSign: r.callSign,
			data:     data,
		})
	}
}

func (r *Radio) getRadiosInRange() []*Radio {
	radios := make([]*Radio, 0)
	for i := range allRadios {
		distance := r.Distance(allRadios[i])
		if r != allRadios[i] && distance <= r.SignalRange && distance <= allRadios[i].SignalRange {
			radios = append(radios, allRadios[i])
		}
	}
	return radios
}

type Receiver interface {
	Receive(message Message)
}

type Speaker struct{}

func (s Speaker) Receive(message Message) {
	fmt.Printf("Received from %v the data %v\n", message.callSign, message.data)
}

func NewRadio(callSign string, antenna Antenna, receiver Receiver) *Radio {
	radio := &Radio{
		callSign: callSign,
		receiver: receiver,
		Antenna:  antenna,
	}
	allRadios = append(allRadios, radio)
	return radio
}
