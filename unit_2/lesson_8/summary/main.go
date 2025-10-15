package main

import "fmt"

func main() {
	const (
		distance   = 236000000000000000 // km
		lightSpeed = 299792             // km/s
		seconds    = distance / lightSpeed
		minutes    = seconds / 60
		hours      = minutes / 60
		days       = hours / 24
		years      = days / 365.2425
	)
	fmt.Printf("Canis Major Dwarf is arguably the closest other galaxy to Earth and it is %1.0f light years away.\n", years)
}
