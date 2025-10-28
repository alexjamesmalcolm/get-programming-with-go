package main

import (
	"fmt"
	"strings"
)

// Coordinate gives a row, column location.
// It can be flexibly used as either Sudoku coordinates (0-8, 0-8) or subregion coordinates (0-2, 0-2).
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
	// initial was the given numbers that make the puzzle unique
	initial Grid
	// Grid reflects the current state of the Sudoku board
	Grid
}

func (g Grid) String() string {
	var s = make([]string, 0, 11)
	horizontalRule := strings.Repeat("-", 11)
	for row := range rows {
		var line = make([]string, 0, 11)
		for column := range columns {
			char := fmt.Sprintf("%d", g[row][column])
			if char == "0" {
				char = "•"
			}
			line = append(line, char)
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

func (g Grid) Get(coord Coordinate) (digit int8, err error) {
	if inBounds(coord) {
		digit = g[coord.row][coord.column]
	} else {
		err = ErrBounds(coord)
	}
	return digit, err
}

func (s *Sudoku) Set(coord Coordinate, digit int8) error {
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
	if !s.validByColumn(coord, digit) {
		errs = append(errs, ErrVertical(digit))
	}
	if !s.validByRow(coord, digit) {
		errs = append(errs, ErrHorizontal(digit))
	}
	if !s.validBySubregion(coord, digit) {
		errs = append(errs, ErrSubregion(digit))
	}
	if s.initial[coord.row][coord.column] != 0 {
		errs = append(errs, ErrInitial(coord))
	}
	if len(errs) > 0 {
		return errs
	}
	s.Grid[coord.row][coord.column] = digit
	return nil
}

// inBounds reports whether the coordinate is within a 9x9 Sudoku grid.
func inBounds(coord Coordinate) bool {
	if coord.row < 0 || coord.row > rows {
		return false
	}
	if coord.column < 0 || coord.column > columns {
		return false
	}
	return true
}

// validDigit tests if the supplied digit is 1 to 9 inclusive.
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// getCoordinatesOfRow returns the nine coordinates that make up the row that the provided
// coordinate resides in.
func getCoordinatesOfRow(coord Coordinate) [9]Coordinate {
	var row [9]Coordinate
	for column := range columns {
		row[column] = Coordinate{coord.row, column}
	}
	return row
}

// getCoordinatesOfColumn returns the nine coordinates that make up the column that the provided
// coordinate resides in.
func getCoordinatesOfColumn(coord Coordinate) [9]Coordinate {
	var column [9]Coordinate
	for row := range rows {
		column[row] = Coordinate{row, coord.column}
	}
	return column
}

// getCoordinatesOfSubregion returns the nine coordinates that make up the subregion that the
// provided coordinate resides in.
func getCoordinatesOfSubregion(coord Coordinate) [9]Coordinate {
	var subregion [9]Coordinate
	// regionCoordinate rows and columns can range 0-2 and refer to the overall subregions
	// coordinate amongst the other subregions.
	var regionCoordinate = Coordinate{
		row:    (coord.row / 3) * 3,
		column: (coord.column / 3) * 3,
	}
	for i := range subregion {
		subregion[i] = Coordinate{
			row:    regionCoordinate.row + i/3,
			column: regionCoordinate.column + i%3,
		}
	}
	return subregion
}

// validByCoordinateGroup reports whether the provided digit can be placed at the Coordinate by checking against the coordinate group
func (g Grid) validByCoordinateGroup(
	coord Coordinate,
	digit int8,
	getCoordinateGroup func(coord Coordinate) [9]Coordinate,
) bool {
	for _, comparisonCoord := range getCoordinateGroup(coord) {
		comparisonDigit, err := g.Get(comparisonCoord)
		if err != nil {
			return false
		}
		if comparisonDigit == digit {
			return false
		}
	}
	return true
}

func (g Grid) validByRow(coord Coordinate, digit int8) bool {
	return g.validByCoordinateGroup(coord, digit, getCoordinatesOfRow)
}

func (g Grid) validByColumn(coord Coordinate, digit int8) bool {
	return g.validByCoordinateGroup(coord, digit, getCoordinatesOfColumn)
}

func (g Grid) validBySubregion(coord Coordinate, digit int8) bool {
	return g.validByCoordinateGroup(coord, digit, getCoordinatesOfSubregion)
}

func NewSudoku(initial Grid) Sudoku {
	current := initial
	return Sudoku{initial, current}
}

type SafeSudokuSetter struct {
	err    error
	sudoku Sudoku
}

func (ss *SafeSudokuSetter) SafeSet(coord Coordinate, digit int8) {
	if ss.err != nil {
		return
	}
	ss.err = ss.sudoku.Set(coord, digit)
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
	s := SafeSudokuSetter{sudoku: NewSudoku(initialGrid)}
	s.SafeSet(Coordinate{5, 2}, 3)
	s.SafeSet(Coordinate{6, 0}, 9)
	s.SafeSet(Coordinate{5, 6}, 8)
	s.SafeSet(Coordinate{2, 6}, 5)
	s.SafeSet(Coordinate{7, 1}, 8)
	s.SafeSet(Coordinate{7, 2}, 7)
	s.SafeSet(Coordinate{1, 1}, 7)
	s.SafeSet(Coordinate{2, 8}, 7)
	s.SafeSet(Coordinate{4, 2}, 6)
	s.SafeSet(Coordinate{7, 6}, 6)
	s.SafeSet(Coordinate{8, 6}, 1)
	s.SafeSet(Coordinate{6, 2}, 1)
	s.SafeSet(Coordinate{2, 0}, 1)
	s.SafeSet(Coordinate{7, 0}, 2)
	s.SafeSet(Coordinate{7, 7}, 3)
	s.SafeSet(Coordinate{6, 8}, 4)
	s.SafeSet(Coordinate{8, 0}, 3)
	s.SafeSet(Coordinate{8, 1}, 4)
	s.SafeSet(Coordinate{8, 2}, 5)
	if s.err != nil {
		fmt.Println(s.err)
		return
	}
	fmt.Println(s.sudoku)
}
