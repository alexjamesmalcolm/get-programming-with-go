package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere) // Prints <nil>
	// fmt.Println(*nowhere) // Panic: nil pointer dereference
	if nowhere != nil {
		fmt.Println(*nowhere) // Can't panic because it won't run if nil
	}
}
