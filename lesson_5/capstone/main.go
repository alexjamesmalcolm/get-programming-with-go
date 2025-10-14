// Capstone: Ticket to Mars
package main

import (
	"fmt"
	"math/rand"
)

type ticket struct {
	spaceline string
	days      int
	tripType  string
	price     int
}

func generateRandomNumberInclusive(lowest int, highest int) int {
	var (
		difference          = highest - lowest
		offsetToBeInclusive = difference + 1
	)
	return rand.Intn(offsetToBeInclusive) + lowest
}

func printTicketTable(tickets []ticket) {
	fmt.Println("Spaceline\t\tDays\tTrip type\tPrice")
	fmt.Println("=============================================")
	for i := range tickets {
		ticket := tickets[i]
		fmt.Println(ticket.spaceline, "\t", ticket.days, "\t", ticket.tripType, "\t", ticket.price)
	}
}

func generateTicket() ticket {
	var spaceline = ""
	switch num := generateRandomNumberInclusive(1, 3); num {
	case 1:
		spaceline = "Space Adventures"
	case 2:
		spaceline = "SpaceX"
	case 3:
		spaceline = "Virgin Galactic"
	}
	isRoundTrip := rand.Intn(2) == 1
	var tripType string
	if isRoundTrip {
		tripType = "Round-trip"
	} else {
		tripType = "One-way"
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
	return ticket{
		spaceline: spaceline,
		tripType:  tripType,
		days:      days,
		price:     price,
	}
}

func main() {
	var tickets = []ticket{}
	for count := 10; count > 0; count-- {
		tickets = append(tickets, generateTicket())
	}
	printTicketTable(tickets)
}
