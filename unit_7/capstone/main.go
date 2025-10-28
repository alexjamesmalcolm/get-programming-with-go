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
	start = command(2)
	stop  = command(3)
)

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandChannel chan command
}

// drive is responsible for driving the rover. It is expected to be started in a goroutine.
func (rd RoverDriver) drive() {
	pos := image.Point{0, 0}
	direction := image.Point{1, 0}
	speed := 0
	updateInterval := 250 * time.Millisecond
	nextMove := time.NewTicker(updateInterval)
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
			case stop:
				speed = 0
			case start:
				speed = 1
			}
			log.Printf("facing %v at speed %v", direction, speed)
		case <-nextMove.C:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("moved to %v", pos)
		}
	}
}

// Left turns the rover left (90° counterclockwise)
func (rd RoverDriver) Left() {
	rd.commandChannel <- left
}

// Right turns the rover right (90° clockwise)
func (rd RoverDriver) Right() {
	rd.commandChannel <- right
}

// Start increases the speed to full
func (rd RoverDriver) Start() {
	rd.commandChannel <- start
}

// Stop brings the rover to a stop
func (rd RoverDriver) Stop() {
	rd.commandChannel <- stop
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{commandChannel: make(chan command)}
	go r.drive()
	return r
}

// MarsGrid represents a grid of some of the surface of Mars.
// It may be used concurrently by different goroutines.
type MarsGrid struct {
	// TODO
}

// Occupy occupies a cell at the given point in the grid. It returns nil if the point is already
// occupied or the point is outside the grid. Otherwise it returns a value that can be used to move
// to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	// TODO
}

func NewMarsGrid(width, height int) *MarsGrid {
	// TODO
}

// Occupier represents an occupied cell in the grid.
// It may be used concurrently by different goroutines.
type Occupier struct {
	// TODO
}

// Move moves the occupier to a different cell in the grid. It reports whether the move was
// successful. It might fail because it was trying to move outside the grid or because the cell
// it's trying to move into is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) Move(p image.Point) bool {
	// TODO
}

func main() {}
