package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

func newLocation(name string, lat, long coordinate) location {
	return location{name, lat.decimal(), long.decimal()}
}

type world struct {
	radius float64
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

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	degrees, minutes, seconds float64
	hemisphere                byte
}

// decimal converts a d/m/s coordinate to a decimal degrees.
func (coord coordinate) decimal() float64 {
	sign := 1.0
	switch coord.hemisphere {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (coord.degrees + coord.minutes/60 + coord.seconds/3600)
}

func triangularNumbers(n int) int {
	n--
	return n * (n + 1) / 2
}

func uniquePairs(locations []location) [][2]location {
	triNum := triangularNumbers(len(locations))
	var pairsOfLocations [][2]location = make([][2]location, 0, triNum)
	for _, a := range locations {
		for _, b := range locations {
			if a == b {
				continue
			}
			var alreadyHasReversePair = false
			for _, pair := range pairsOfLocations {
				if pair[0] == b && pair[1] == a {
					alreadyHasReversePair = true
					break
				}
			}
			if alreadyHasReversePair {
				continue
			}
			var pair = [2]location{a, b}
			pairsOfLocations = append(pairsOfLocations, pair)
		}
	}
	return pairsOfLocations
}

func printDistance(w world, a, b location) {
	distance := w.distance(a, b)
	fmt.Printf("%v to %v is %.3f km\n", a.name, b.name, distance)
}

func main() {
	mountSharp := newLocation("Mount Sharp", coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newLocation("Olympus Mons", coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	marsLocations := []location{
		newLocation("Columbia Memorial Station", coordinate{14, 35, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'}),
		newLocation("Challenger Memorial Station", coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'}),
		newLocation("Bradbury Landing", coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'}),
		newLocation("Elysium Planitia", coordinate{4, 30, 0, 'S'}, coordinate{135, 54, 0, 'E'}),
		mountSharp,
		olympusMons,
	}
	for _, l := range marsLocations {
		fmt.Printf("%v: lat=%0.3f, long=%0.3f\n", l.name, l.lat, l.long)
	}

	mars := world{3389.5}

	for _, pair := range uniquePairs(marsLocations) {
		printDistance(mars, pair[0], pair[1])
	}

	earth := world{6371}

	london := newLocation("London, England", coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation("Paris, France", coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	printDistance(earth, london, paris)

	columbus := newLocation("Columbus, OH", coordinate{39, 57, 44, 'N'}, coordinate{83, 0, 2, 'W'})
	washington := newLocation("Washington, DC", coordinate{38, 54, 17, 'N'}, coordinate{77, 0, 59, 'W'})
	printDistance(earth, columbus, washington)

	printDistance(mars, mountSharp, olympusMons)

}
