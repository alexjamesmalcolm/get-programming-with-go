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

func numberOneToTen() int {
	// return rand.Intn(10) + 1
	return generateRandomNumberInclusive(1, 10)
}

func main() {
	var num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
	num = numberOneToTen()
	fmt.Println(num)
}
