// For this challenge, you will build a simulation of underpopulation, overpopulation, and
// reproduction called Conway's Game of Life (see https://mng.bz/xOyY). The simulation is played
// out on a two-dimensional grid of cells. As such, this challenge focuses on slices.

// Each cell has eight adjacent cells in the horizontal, vertical, and diagonal directions. In each
// generation, cells live or die based on the number of living neighbors.
package main

import (
	"fmt"
	"math/rand"
)

const (
	// width of the universe
	width = 80
	// height of the universe
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	var universe Universe
	for range width {
		column := make([]bool, height)
		universe = append(universe, column)
	}
	return universe
}

func (u Universe) Show() {
	const (
		liveCell = '*'
		deadCell = ' '
	)
	for y := range height {
		for x := range width {
			if u[x][y] {
				fmt.Printf("%c", liveCell)
			} else {
				fmt.Printf("%c", deadCell)
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for y := range height {
		for x := range width {
			if rand.Intn(4) == 0 {
				u[x][y] = true
			} else {
				u[x][y] = false
			}
		}
	}
}

func main() {
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
}
