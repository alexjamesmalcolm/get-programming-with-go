package location

import (
	"capstone/radio"
	"fmt"
	"math"
	"math/rand"
)

type Planet struct {
	radius float64
}
type Coordinate struct {
	Lat, Long float64
}
type Location struct {
	Coordinate
	Planet
	OnSurface bool
}

var (
	Earth Planet = Planet{6371}
	Mars  Planet = Planet{3389.5}
)

// rad converts degrees to radians.
func rad(degrees float64) (radians float64) {
	return degrees * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines.
func surfaceDistance(w Planet, p1, p2 Location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

var distanceBetweenEarthAndMars = generateRandomEarthMarsDistance()

func (l Location) Distance(other radio.Distancer) float64 {
	if otherLocation, ok := other.(Location); ok {
		if l.Planet == otherLocation.Planet {
			if l.OnSurface && otherLocation.OnSurface {
				return surfaceDistance(l.Planet, l, otherLocation)
			}
			return float64(generateRandomMarsOrbitDistance())
		}
		return float64(distanceBetweenEarthAndMars)
	}
	panic(fmt.Sprintf("unable to determine distance between %v and %v", l, other))
}

func generateRandomNumberInclusive(lowest int, highest int) int {
	var (
		difference          = highest - lowest
		offsetToBeInclusive = difference + 1
	)
	return rand.Intn(offsetToBeInclusive) + lowest
}

func generateRandomEarthMarsDistance() uint {
	// Closest 56,000,000 and the farthest is 401,000,000 km
	const closest = 56000000   // 56,000,000 km
	const farthest = 401000000 // 401,000,000 km
	return uint(generateRandomNumberInclusive(closest, farthest))
}

func generateRandomMarsOrbitDistance() int {
	// https://en.wikipedia.org/wiki/MAVEN
	const (
		closest  = 180
		farthest = 4500
	)
	return generateRandomNumberInclusive(closest, farthest)
}
