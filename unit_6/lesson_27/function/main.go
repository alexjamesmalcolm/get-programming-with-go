package main

import "fmt"

func callFunction(fn func(a, b int) int) int {
	return fn(1, 2) // Without a guard this will panic
}

func main() {
	var fn func(a, b int) int // Is initialized with a nil value
	fnPointer := &fn
	fmt.Println(fn == nil) // Prints true
	fmt.Println(fnPointer)
	callFunction(fn)

}
