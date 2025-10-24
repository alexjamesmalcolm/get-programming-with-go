package main

import "fmt"

func main() {
	answer := 42
	// In &answer the & is the "address operator"
	address := &answer
	fmt.Println(address) // Prints some memory address
	// The * "dereferences" the address back into the actual value
	fmt.Println(*address) // Prints 42
}
