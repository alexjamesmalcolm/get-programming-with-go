package main

import "fmt"

func conciseSwitch() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"
	switch command {
	case "go east": // This is a matching switch statement, there is no conditional just a value
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}

func fallThroughSwitch() {
	var room = "lake"
	switch {
	case room == "cave": // This is an expression switch statement
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}

func modifiedFallThroughSwitch() {
	var room = "lake"

	switch room {
	case "cave": // This is an expression switch statement
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}

func main() {
	// conciseSwitch()
	// fallThroughSwitch()
	modifiedFallThroughSwitch()
}
