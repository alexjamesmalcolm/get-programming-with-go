package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	left  = command(0)
	right = command(1)
)

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandChannel chan command
}

func (rd RoverDriver) drive() {
	pos := image.Point{0, 0}
	direction := image.Point{1, 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-rd.commandChannel:
			switch c {
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{commandChannel: make(chan command)}
	go r.drive()
	return r
}

func main() {
	NewRoverDriver()
	time.Sleep(10 * time.Second)
}
