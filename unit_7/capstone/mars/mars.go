// Package mars contains all the location data for navigating the surface of Mars
package mars

import (
	"fmt"
	"image"
	"math/rand"
	"sync"
)

// MarsGrid represents a grid of some of the surface of Mars.
// It may be used concurrently by different goroutines.
type MarsGrid struct {
	// mu guards the cells grid
	mu    sync.RWMutex
	cells [][]cell
}

func (g *MarsGrid) Width() int {
	return len(g.cells)
}
func (g *MarsGrid) Height() int {
	return len(g.cells[0])
}

func (g *MarsGrid) PrintOccupiers() {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for y := range g.Height() {
		for x := range g.Width() {
			if g.cells[x][y].occupier == nil {
				fmt.Printf(".")
			} else {
				fmt.Printf("X")
			}
		}
		fmt.Println()
	}
}
func (g *MarsGrid) PrintLifeSigns() {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for y := range g.Height() {
		for x := range g.Width() {
			fmt.Printf("%4d", g.cells[x][y].groundData.LifeSigns)
		}
		fmt.Println()
	}
}

type SensorData struct {
	LifeSigns int
}

type cell struct {
	groundData SensorData
	occupier   *Occupier
}

// Occupy occupies a cell at the given point in the grid. It returns nil if the point is already
// occupied or the point is outside the grid. Otherwise it returns a value that can be used to move
// to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	if !g.isInBounds(p) {
		return nil
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.cells[p.X][p.Y].occupier != nil {
		return nil
	}
	g.cells[p.X][p.Y].occupier = &Occupier{grid: g}
	return g.cells[p.X][p.Y].occupier
}

func (g *MarsGrid) move(o *Occupier, to image.Point) bool {
	if o == nil || !g.isInBounds(to) {
		return false
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.cells[to.X][to.Y].occupier != nil {
		return false
	}
	for x := range g.Width() {
		for y := range g.Height() {
			if g.cells[x][y].occupier == o {
				g.cells[x][y].occupier = nil
			}
		}
	}
	g.cells[to.X][to.Y].occupier = o
	return true
}

func (g *MarsGrid) isInBounds(p image.Point) bool {
	if p.X < 0 || p.X >= g.Width() {
		return false
	}
	if p.Y < 0 || p.Y >= g.Height() {
		return false
	}
	return true
}

func NewMarsGrid(width, height int) *MarsGrid {
	var cells = make([][]cell, 0, width)
	for x := range width {
		cells = append(cells, make([]cell, height))
		for y := range height {
			cells[x][y] = cell{
				groundData: SensorData{LifeSigns: rand.Intn(1001)},
			}
		}
	}
	return &MarsGrid{
		cells: cells,
	}
}

// Occupier represents an occupied cell in the grid.
// It may be used concurrently by different goroutines.
type Occupier struct {
	// Location image.Point
	grid *MarsGrid
}

// Move moves the occupier to a different cell in the grid. It reports whether the move was
// successful. It might fail because it was trying to move outside the grid or because the cell
// it's trying to move into is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) Move(p image.Point) bool {
	return o.grid.move(o, p)
}
