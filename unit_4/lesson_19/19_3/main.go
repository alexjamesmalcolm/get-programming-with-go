package main

import (
	"fmt"
	"math"
)

func main() {
	// temperature := make(map[float64]int, 8)

	temperatures := []float64{-28, 32, -31, -29, -23, -29, -28, -33}
	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
	// frequency := make(map[float64]int)

	// for _, t := range temperatures {
	// 	frequency[t]++
	// }
	// for t, num := range frequency {
	// 	fmt.Printf("%+.2f occurs %d times\n", t, num)
	// }
}
