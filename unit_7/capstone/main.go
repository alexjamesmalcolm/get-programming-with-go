package main

import (
	"capstone/mars"
	"capstone/rover"
	"image"
)

func main() {
	grid := mars.NewMarsGrid(10, 10)
	driver1 := rover.NewRoverDriver(grid, image.Point{1, 5})
	driver2 := rover.NewRoverDriver(grid, image.Point{1, 7})
	driver1.Start()
	driver2.Start()
	for {
		select {}
	}

}
