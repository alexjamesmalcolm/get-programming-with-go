package main

import "fmt"

func main() {
	const speed = 100800   // km/h
	const hoursInADay = 24 // hours
	var (
		distance = 96300000 // km
		hours    = distance / speed
		days     = hours / hoursInADay
	)
	fmt.Println(
		"It would take the SpaceX Interplanetary Transport System",
		days, "days to get to Mars from Earth",
	)
}
