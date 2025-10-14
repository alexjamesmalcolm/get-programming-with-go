// Capstone: Ticket to Mars
package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumberInclusive(lowest int, highest int) int {
	var (
		difference          = highest - lowest
		offsetToBeInclusive = difference + 1
	)
	return rand.Intn(offsetToBeInclusive) + lowest
}

func main() {
	fmt.Println("Spaceline\t\tDays\tTrip type\tPrice")
	fmt.Println("=====================================================")
	for count := 10; count > 0; count-- {
		var spaceline = ""
		switch num := generateRandomNumberInclusive(1, 3); num {
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "SpaceX\t\t"
		case 3:
			spaceline = "Virgin Galactic\t"
		}
		isRoundTrip := rand.Intn(2) == 1
		var tripType string
		if isRoundTrip {
			tripType = "Round-trip"
		} else {
			tripType = "One-way\t"
		}
		var (
			// km/s
			speed = generateRandomNumberInclusive(16, 30)
			// Cost of the ticket in millions of dollars
			price = speed + 20
			// Distance between the Earth and Mars in km
			distance = 62100000
			// Number of days for a one way trip
			days = distance / (speed * 86400)
		)
		if isRoundTrip {
			price *= 2
		}
		fmt.Print(spaceline, "\t", days, "\t", tripType, "\t$ ", price, "\n")
	}
}
