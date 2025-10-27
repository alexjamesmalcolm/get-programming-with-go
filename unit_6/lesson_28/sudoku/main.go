package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("outside of grid boundaries")
	ErrDigit  = errors.New("invalid digit")
)

// Grid is a Sudoku grid
type Grid [rows][columns]int8

func (g Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return ErrDigit
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func main() {
	var g Grid
	err := g.Set(0, 0, 10)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}
	fmt.Println(g)
}
