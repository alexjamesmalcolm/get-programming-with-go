package main

import "fmt"

func quickCheck() {
	const speed = 100800    // km/h
	const hoursInADay = 24  // hours
	var distance = 96300000 // km
	var hours = distance / speed
	var days = hours / hoursInADay
	fmt.Println(
		"It would take the SpaceX Interplanetary Transport System",
		days, "days to get to Mars from Earth",
	)
}

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
	quickCheck()
}
