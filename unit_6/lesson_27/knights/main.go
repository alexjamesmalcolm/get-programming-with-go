package main

import "fmt"

type item string
type character struct {
	name     string
	leftHand *item
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("It's %v and he's not holding anything.", c.name)
	}
	return fmt.Sprintf("It's %v with a %v in his left hand.", c.name, *c.leftHand)
}
func (c *character) pickUp(i *item) {
	if i == nil {
		fmt.Printf("There's nothing for %v to pick up.\n", c.name)
		return
	}
	fmt.Printf("%v has picked up %v with his left hand.\n", c.name, *i)
	c.leftHand = i
}
func (c *character) give(to *character) {
	item := c.leftHand
	if item == nil {
		fmt.Printf("%v has nothing in his hand to give %v.\n", c.name, to.name)
		return
	}
	c.leftHand = nil
	to.leftHand = item
	fmt.Printf("%v has given %v to %v.\n", c.name, *item, to.name)
}

func main() {
	knight := character{name: "The knight who says nil"}
	arthur := character{name: "Arthur"}
	arthur.give(&knight)

	var nothing *item
	arthur.pickUp(nothing)
	var shrubbery item = "Good shrubbery"
	arthur.pickUp(&shrubbery)

	arthur.give(&knight)
}
