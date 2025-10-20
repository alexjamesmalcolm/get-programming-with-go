package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)

	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity) // Prints the struct with its fields names
}
