package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)              // Prints 0.3333333333333333
	fmt.Printf("%v\n", third)       // Prints 0.3333333333333333
	fmt.Printf("%f\n", third)       // Prints 0.333333
	fmt.Printf("%0.3f\n", third)    // Prints 0.333
	fmt.Printf("%4.2f\n", third)    // Prints 0.33 because it has a width of 4 and a precision of 2
	fmt.Printf("%05.2f\n", third)   // Prints 00.33 because it has a width of 5 and a precision of 2 and it is using zeros for left padding
	fmt.Printf("%09.4f\n", 15.1021) // Prints 0015.1021
}
