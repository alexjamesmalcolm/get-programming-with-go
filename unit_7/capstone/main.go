package main

import (
	"capstone/mars"
	"capstone/rover"
	"image"
)

func main() {
	grid := mars.NewMarsGrid(100, 100)
	driver1 := rover.NewRoverDriver("1", grid, image.Point{1, 1})
	driver2 := rover.NewRoverDriver("2", grid, image.Point{1, 3})
	driver3 := rover.NewRoverDriver("3", grid, image.Point{1, 5})
	driver4 := rover.NewRoverDriver("4", grid, image.Point{1, 7})
	driver5 := rover.NewRoverDriver("5", grid, image.Point{1, 9})
	driver1.Start()
	driver2.Start()
	driver3.Start()
	driver4.Start()
	driver5.Start()
	for {
		select {}
	}
}
