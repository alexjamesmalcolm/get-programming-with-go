package main

import "fmt"

func panicHandler() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}

func main() {
	defer panicHandler()
	panic("I forgot my towel")
	fmt.Println("You can borrow mine.") // This code is unreachable
}
