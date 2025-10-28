package main

import "image"

// MarsGrid represents a grid of some of the surface of Mars.
// It may be used concurrently by different goroutines.
type MarsGrid struct {
	// TODO
}

// Occupy occupies a cell at the given point in the grid. It returns nil if the point is already
// occupied or the point is outside the grid. Otherwise it returns a value that can be used to move
// to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier

// Occupier represents an occupied cell in the grid.
// It may be used concurrently by different goroutines.
type Occupier struct {
	// TODO
}

// Move moves the occupier to a different cell in the grid. It reports whether the move was
// successful. It might fail because it was trying to move outside the grid or because the cell
// it's trying to move into is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) Move(p image.Point) bool

func main() {}
