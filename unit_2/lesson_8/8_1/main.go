package main

import "fmt"

func main() {
	const (
		lightSpeed    = 299792 // km/s
		secondsPerDay = 86400
	)
	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("This is", days, "days of travel at light speed.")

	const (
		closestEarthMarsDistance  uint64 = 56e6
		farthestEarthMarsDistance uint64 = 401e6
	)
	// fmt.Println("The distance between Mars and Earth ranges from", closestEarthMarsDistance, "km to", farthestEarthMarsDistance, "km.")
	fmt.Printf("The distance between Mars and Earth ranges from %d km to %d km.\n", closestEarthMarsDistance, farthestEarthMarsDistance)
}
