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

// Alive wraps around the edges of the universe so the top is connected to the bottom and the left
// is connected to the right at the edges.
func (u Universe) Alive(x, y int) bool {
	fmt.Println("Alive received", x, y)
	if x < 0 {
		x += width
	} else {
		x = x % width
	}
	if y < 0 {
		y += height
	} else {
		y = y % height
	}
	return u[x][y]
}

// Neighbors counts the number of live neighbors for a given cell, from 0 to 8.
func (u Universe) Neighbors(x, y int) uint8 {
	var living uint8 = 0
	livingSlice := []bool{
		u.Alive(x-1, y-1),
		u.Alive(x, y-1),
		u.Alive(x+1, y-1),
		u.Alive(x-1, y),
		u.Alive(x+1, y),
		u.Alive(x-1, y-1),
		u.Alive(x, y-1),
		u.Alive(x+1, y-1),
	}
	for _, isAlive := range livingSlice {
		if isAlive {
			living++
		}
	}
	return living
}

func main() {
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
	fmt.Println(universe.Neighbors(10, 10))
	fmt.Println(universe.Neighbors(0, 0))
	fmt.Println(universe.Neighbors(80, 15))
	fmt.Println(universe.Neighbors(200, 200))
}
