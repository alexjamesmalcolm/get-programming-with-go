// Package mars contains all the location data for navigating the surface of Mars
package mars

import (
	"fmt"
	"image"
	"math/rand"
	"strings"
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

func (g *MarsGrid) StringOccupiers() string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	var lines []string
	for y := range g.Height() {
		var line []string
		for x := range g.Width() {
			if g.cells[x][y].occupier == nil {
				line = append(line, ".")
			} else {
				line = append(line, "X")
			}
		}
		lines = append(lines, strings.Join(line, ""))
	}
	return strings.Join(lines, "\n")
}
func (g *MarsGrid) PrintOccupiers() {
	fmt.Println(g.StringOccupiers())
}

func (g *MarsGrid) StringLifeSigns() string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	var lines []string
	for y := range g.Height() {
		var line []string
		for x := range g.Width() {
			line = append(line, fmt.Sprintf("%4d", g.cells[x][y].groundData.LifeSigns))
		}
		lines = append(lines, strings.Join(line, ""))
	}
	return strings.Join(lines, "\n")
}
func (g *MarsGrid) PrintLifeSigns() {
	fmt.Println(g.StringLifeSigns())
}

type SensorData struct {
	Location  image.Point
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
	g.cells[p.X][p.Y].occupier = &Occupier{grid: g, location: p}
	return g.cells[p.X][p.Y].occupier
}

func (g *MarsGrid) move(from image.Point, to image.Point) bool {
	if !g.isInBounds(to) {
		return false
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.cells[to.X][to.Y].occupier != nil {
		return false
	}
	if g.cells[from.X][from.Y].occupier == nil {
		return false
	}
	g.cells[to.X][to.Y].occupier = g.cells[from.X][from.Y].occupier
	g.cells[from.X][from.Y].occupier = nil
	return true
}

func (g *MarsGrid) getSensorData(p image.Point) SensorData {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.cells[p.X][p.Y].groundData
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
				groundData: SensorData{
					LifeSigns: rand.Intn(1001),
					Location:  image.Point{x, y},
				},
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
	location image.Point
	grid     *MarsGrid
}

func (o *Occupier) Point() image.Point {
	return o.location
}

// Move moves the occupier to a different cell in the grid. It reports whether the move was
// successful. It might fail because it was trying to move outside the grid or because the cell
// it's trying to move into is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) Move(to image.Point) bool {
	didMove := o.grid.move(o.location, to)
	if didMove {
		o.location = to
	}
	return didMove
}

func (o *Occupier) GetSensorData() SensorData {
	return o.grid.getSensorData(o.location)
}
