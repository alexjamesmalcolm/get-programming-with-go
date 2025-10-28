package main

import (
	"capstone/mars"
	"fmt"
	"image"
)

func main() {
	grid := mars.NewMarsGrid(4, 3)
	grid.PrintLifeSigns()
	grid.PrintOccupiers()
	occupier := grid.Occupy(image.Point{1, 1})
	if occupier == nil {
		fmt.Println("Expected occupier but got nil")
		return
	}
	fmt.Println()
	grid.PrintOccupiers()

	secondOccupier := grid.Occupy(image.Point{1, 2})
	if secondOccupier == nil {
		fmt.Println("Expected second occupier but got nil")
	}
	fmt.Println()
	grid.PrintOccupiers()

	didMove := occupier.Move(image.Point{2, 2})

	if didMove {
		fmt.Println("We were able to move the original occupier to (2, 2)")
	} else {
		fmt.Println("We were not able to move the original occupier")
	}
	grid.PrintOccupiers()

	// secondOccupier = grid.Occupy(image.Point{2, 2})
	// fmt.Println("Expecting to be nil:", secondOccupier)

}
