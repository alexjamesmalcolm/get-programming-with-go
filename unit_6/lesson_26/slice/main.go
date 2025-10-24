package main

import "fmt"

// reclassify takes a pointer to a slice of strings because it is going to mutate that slice by
// reducing its length.
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets) // What happened to Pluto?
}
