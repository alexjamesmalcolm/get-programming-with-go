package main

import (
	"fmt"
	"math"
)

type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type world struct {
	radius float64
}
type location struct {
	name      string
	lat, long float64
}

type celsius float64

func (l location) description() string {
	return fmt.Sprintf("%v: (%.4f, %.4f)", l.name, l.lat, l.long)
}

// rad converts degrees to radians.
func rad(degrees float64) (radians float64) {
	return degrees * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type gps struct {
	current     location
	destination location
	world       world
}

func (gps gps) distance() float64 {
	return gps.world.distance(gps.current, gps.destination)
}

func (gps gps) message() string {
	return fmt.Sprintf("%.2f km remaining until %v", gps.distance(), gps.destination.name)
}

type rover struct {
	gps
}

func main() {
	bradbury := location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	t := temperature{high: -1, low: -78}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("A balmy %v° C\n", report.temperature.high)
	fmt.Printf("Average %v° C\n", report.temperature.average())

	mars := world{3389.5}
	gps := gps{
		current:     bradbury,
		destination: location{name: "Elysium Planitia", lat: 4.5, long: 135.9},
		world:       mars,
	}
	curiosity := rover{gps}
	fmt.Printf("Curiosity currently at %v.\n", curiosity.current.description())
	fmt.Println(curiosity.message())
}
