package main

import (
	"fmt"
	"strings"
)

type Coordinate struct{ row, column int }

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

type ErrVertical int8

func (e ErrVertical) Error() string {
	return fmt.Sprintf("%d already exists in column", e)
}

type ErrHorizontal int8

func (e ErrHorizontal) Error() string {
	return fmt.Sprintf("%d already exists in row", e)
}

type ErrSubregion int8

func (e ErrSubregion) Error() string {
	return fmt.Sprintf("%d already exists in 3x3 subregion", e)
}

type ErrBounds Coordinate

func (e ErrBounds) Error() string {
	return fmt.Sprintf("(%v, %v) is outside of the grid boundaries", e.row, e.column)
}

type ErrDigit int8

func (e ErrDigit) Error() string {
	return fmt.Sprintf("%d is not a valid digit as it is not 1-9", e)
}

type ErrInitial Coordinate

func (e ErrInitial) Error() string {
	return fmt.Sprintf("(%v, %v) cannot be changed because it was given", e.row, e.column)
}

const rows, columns = 9, 9

type Grid [rows][columns]int8
type Sudoku struct {
	initial Grid
	current Grid
}

func (sudoku Sudoku) String() string {
	return sudoku.current.String()
}

func (grid Grid) String() string {
	var s = make([]string, 0, 11)
	horizontalRule := strings.Repeat("-", 11)
	for row := range rows {
		var line = make([]string, 0, 11)
		for column := range columns {
			line = append(line, fmt.Sprintf("%d", grid[row][column]))
			if (column+1)%3 == 0 && column+1 != columns {
				line = append(line, "|")
			}
		}
		s = append(s, strings.Join(line, ""))
		if (row+1)%3 == 0 && row+1 != rows {
			s = append(s, horizontalRule)
		}
	}
	return strings.Join(s, "\n")
}

func (grid Grid) Get(coord Coordinate) (digit int8, err error) {
	if inBounds(coord) {
		digit = grid[coord.row][coord.column]
	} else {
		err = ErrBounds(coord)
	}
	return digit, err
}
func (sudoku Sudoku) Get(coord Coordinate) (digit int8, err error) {
	return sudoku.current.Get(coord)
}

func (sudoku *Sudoku) Set(coord Coordinate, digit int8) error {
	var errs SudokuError
	if !inBounds(coord) {
		errs = append(errs, ErrBounds(coord))
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit(digit))
	}
	if len(errs) > 0 {
		return errs
	}
	if !sudoku.current.validByColumn(coord, digit) {
		errs = append(errs, ErrVertical(digit))
	}
	if !sudoku.current.validByRow(coord, digit) {
		errs = append(errs, ErrHorizontal(digit))
	}
	if !sudoku.current.validBySubregion(coord, digit) {
		errs = append(errs, ErrSubregion(digit))
	}
	if sudoku.initial[coord.row][coord.column] != 0 {
		errs = append(errs, ErrInitial(coord))
	}
	if len(errs) > 0 {
		return errs
	}
	sudoku.current[coord.row][coord.column] = digit
	return nil
}

func inBounds(coord Coordinate) bool {
	if coord.row < 0 || coord.row > rows {
		return false
	}
	if coord.column < 0 || coord.column > columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func getCoordinatesOfRow(coord Coordinate) (row [9]Coordinate) {
	for column := range columns {
		row[column] = Coordinate{coord.row, column}
	}
	return row
}
func getCoordinatesOfColumn(coord Coordinate) (column [9]Coordinate) {
	for row := range rows {
		column[row] = Coordinate{row, coord.column}
	}
	return column
}
func getCoordinatesOfSubregion(coord Coordinate) (subregion [9]Coordinate) {
	return subregion
}

func (grid Grid) validByConstraintGroup(
	coord Coordinate,
	digit int8,
	getter func(coord Coordinate) [9]Coordinate,
) bool {
	for _, comparisonCoord := range getter(coord) {
		comparisonDigit, err := grid.Get(comparisonCoord)
		if err != nil {
			return false
		}
		if comparisonDigit == digit {
			return false
		}
	}
	return true
}

func (grid Grid) validByRow(coord Coordinate, digit int8) bool {
	return grid.validByConstraintGroup(coord, digit, getCoordinatesOfRow)
}

func (grid Grid) validByColumn(coord Coordinate, digit int8) bool {
	return grid.validByConstraintGroup(coord, digit, getCoordinatesOfColumn)
}

func (grid Grid) validBySubregion(coord Coordinate, digit int8) bool {
	return grid.validByConstraintGroup(coord, digit, getCoordinatesOfSubregion)
}

func NewSudoku(initial Grid) Sudoku {
	current := initial
	return Sudoku{initial, current}
}

func main() {
	initialGrid := Grid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	s := NewSudoku(initialGrid)
	s.Set(Coordinate{0, 1}, 3)
	fmt.Println(s)
}
