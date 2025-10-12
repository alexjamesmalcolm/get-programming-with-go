package main

import "fmt"

func main() {
	const marsGravityPercentageOfEarth = 0.3783
	var weight = 207.0 // lbs
	// weight = weight * marsGravityPercentageOfEarth
	weight *= marsGravityPercentageOfEarth
	fmt.Println("My weight on Mars would be", weight, "lbs")
}
