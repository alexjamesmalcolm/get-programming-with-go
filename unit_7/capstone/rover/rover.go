package rover

import (
	"capstone/location"
	"capstone/mars"
	"capstone/radio"
	"image"
	"log"
)

type Rover struct {
	name   string
	driver *RoverDriver
	radio  *radio.Radio
}

func (r *Rover) command() {
	for {
		select {
		case data := <-r.driver.data:
			log.Printf("%v - detected signs of life, sending to Earth", r.name)
			r.radio.SendData("Earth", data)
		}
	}
}

func NewRover(name string, g *mars.MarsGrid, p image.Point, l location.Location) (*Rover, error) {
	driver, err := NewRoverDriver(name, g, p)
	if err != nil {
		return nil, err
	}
	r := &Rover{
		name:   name,
		driver: driver,
		radio: radio.NewRadio(name, &radio.Antenna{
			DataRate:    250,
			SignalRange: 5000,
			Distancer:   l,
		}),
	}
	go r.command()
	return r, nil
}
