package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{-28, 32, -31, -29, -23, -29, -28, -33}
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}
	fmt.Println("Map of temperatures", temperatures)
	fmt.Println("Set of temperatures", set)

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println("Sorted unique temperatures", unique)
}
