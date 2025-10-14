package main

import (
	"fmt"
	"math/rand"
)

func uncondensedForLoop() {
	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println("Uncondensed Loop - Count in the loop:", count)
	}
	fmt.Println("Uncondensed Loop - Count after the loop:", count)
}

func shortLoop() {
	for count := 10; count > 0; count-- {
		fmt.Println("Short Loop - Count in the loop:", count)
	}
	// Cannot access "count" anymore because it was scoped to the loop
}

func shortIfStatement() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
	// Cannot access "num" anymore because it was scoped to the if statement
}

func shortSwitch() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}

func main() {
	uncondensedForLoop()
	shortLoop()
	shortIfStatement()
	shortSwitch()
}
