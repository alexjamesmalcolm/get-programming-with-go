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

func generateRandomEarthMarsDistance() int {
	// Closest 56,000,000 and the farthest is 401,000,000 km
	const closest = 56000000   // 56,000,000 km
	const farthest = 401000000 // 401,000,000 km
	return generateRandomNumberInclusive(closest, farthest)
}

func main() {
	var distance = generateRandomEarthMarsDistance()
	fmt.Println("A random possible distance between the Earth and Mars could be", distance, "km.")
}
