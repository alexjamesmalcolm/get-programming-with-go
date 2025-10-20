package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury // Assignment creates a copy

	curiosity.long += 0.0106 // Changes to curiosity's location does not also change Bradbury Landing's location
	fmt.Println(bradbury, curiosity)
}
