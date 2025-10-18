package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	giants := planets[4:]
	gas := giants[:2]
	ice := giants[2:]
	fmt.Println(giants, gas, ice)
	fmt.Println()

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII, ice)

	sort.StringSlice(planets[:]).Sort()
	// This does the same thing as the above statement but shorter
	sort.Strings(planets[:])
	fmt.Println(planets) // All the planets are sorted
}
