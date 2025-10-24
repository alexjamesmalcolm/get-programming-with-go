package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	rebeccaPointer := &rebecca
	anotherRebecca := *rebeccaPointer // Makes a copy of Rebecca
	birthday(rebeccaPointer)
	fmt.Printf("%+v\n", rebecca)        // Shows that she is now 15 years old
	fmt.Printf("%+v\n", rebeccaPointer) // Shows that she is now 15 years old
	fmt.Printf("%+v\n", anotherRebecca) // Shows that she is now 15 years old
}
