package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := Planets{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	planets[3:4].terraform()
	planets[6:].terraform()
	fmt.Println(planets)
}
