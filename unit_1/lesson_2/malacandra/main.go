// Malacandra is much nearer than that: we shall make it in about twenty-eight days. - C.S. Lewis, Out of the Silent Planet
package main

import "fmt"

func main() {
	const (
		distance      = 56000000                   // km
		daysOfTravel  = 28                         // days
		hoursInADay   = 24                         // hours
		hoursOfTravel = daysOfTravel * hoursInADay // hours
	)
	fmt.Println("To travel for", daysOfTravel, "days would add up to", hoursOfTravel, "hours of travel.")
	const speed = distance / hoursOfTravel // km/h
	fmt.Println("In order to travel", distance, "km in", daysOfTravel, "days you would have to be moving at", speed, "km/h.")
}
