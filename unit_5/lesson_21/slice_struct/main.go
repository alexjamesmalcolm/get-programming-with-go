package main

import "fmt"

type location struct {
	name      string
	lat, long float64
}

func main() {
	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	fmt.Println(locations)
}
