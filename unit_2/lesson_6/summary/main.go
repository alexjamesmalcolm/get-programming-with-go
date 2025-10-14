package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for piggyBank := 0.0; piggyBank < 20.0; {
		var addition float64
		switch num := rand.IntN(3); num {
		case 0:
			// nickel
			addition = 0.05
		case 1:
			// dime
			addition = 0.10
		case 2:
			// quarter
			addition = 0.25
		}
		fmt.Printf("Total $%05.2f after adding $%3.2f to piggy bank of $%05.2f\n", piggyBank+addition, addition, piggyBank)
		piggyBank += addition
	}
}
