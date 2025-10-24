package main

import "fmt"

type location struct {
	x, y int
}

type turtle struct {
	name string
	location
}

func (t *turtle) moveUp()    { t.y-- }
func (t *turtle) moveDown()  { t.y++ }
func (t *turtle) moveLeft()  { t.x-- }
func (t *turtle) moveRight() { t.x++ }

func main() {
	t := turtle{
		name:     "George",
		location: location{10, 10},
	}
	fmt.Printf("%+v\n", t)
	t.moveDown()
	t.moveDown()
	t.moveLeft()
	fmt.Printf("%+v\n", t)
}
