package rover

import (
	"capstone/mars"
	"image"
	"log"
	"math/rand"
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
	name           string
	commandChannel chan command
	location       *mars.Occupier
}

// drive is responsible for driving the rover. It is expected to be started in a goroutine.
func (rd *RoverDriver) drive() {
	direction := image.Point{1, 0}
	isMoving := false
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
				isMoving = false
			case start:
				isMoving = true
			}
			if isMoving {
				log.Printf("%v facing %v and moving", rd.name, direction)
			} else {
				log.Printf("%v facing %v and stopped", rd.name, direction)
			}
		case <-nextMove.C:
			if isMoving {
				wasAbleToMove := rd.location.Move(rd.location.Point().Add(direction))
				if !wasAbleToMove {
					if rand.Intn(2) == 0 {
						log.Printf("%v got stuck so turning left", rd.name)
						go rd.Left()
					} else {
						log.Printf("%v got stuck so turning right", rd.name)
						go rd.Right()
					}
				} else {
					log.Printf("%v moved to %v", rd.name, rd.location.Point())
				}
			}
		}
	}
}

// Left turns the rover left (90° counterclockwise)
func (rd *RoverDriver) Left() {
	rd.commandChannel <- left
}

// Right turns the rover right (90° clockwise)
func (rd *RoverDriver) Right() {
	rd.commandChannel <- right
}

// Start increases the speed to full
func (rd *RoverDriver) Start() {
	rd.commandChannel <- start
}

// Stop brings the rover to a stop
func (rd *RoverDriver) Stop() {
	rd.commandChannel <- stop
}

func NewRoverDriver(name string, g *mars.MarsGrid, p image.Point) *RoverDriver {
	occupier := g.Occupy(p)
	if occupier == nil {
		return nil
	}
	r := &RoverDriver{
		name:           name,
		commandChannel: make(chan command),
		location:       occupier,
	}
	go r.drive()
	return r
}
