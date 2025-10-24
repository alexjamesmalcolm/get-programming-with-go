package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	// (*timmy).superpower = "flying" // Alternative

	timmy.superpower = "flying" // Automatically dereferenced, no need for (*timmy).superpower
	fmt.Printf("%+v\n", timmy)

	superpowers := &[...]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0]) // Prints flight
	fmt.Println(superpowers[1:2])
}
