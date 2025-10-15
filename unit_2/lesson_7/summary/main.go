package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var piggyBankCents uint16 = 0
	for piggyBankCents < 2000 {
		var additionalCents uint16
		switch num := rand.IntN(3); num {
		case 0:
			// nickel
			additionalCents = 5
		case 1:
			// dime
			additionalCents = 10
		case 2:
			// quarter
			additionalCents = 25
		}
		piggyBankCents += additionalCents
		fmt.Printf("Total $%d.%02d\n", piggyBankCents/100, piggyBankCents%100)
	}
}
