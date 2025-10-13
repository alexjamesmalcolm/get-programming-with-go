package main

import (
	"fmt"
	"math/rand"
)

// Determines if a year is a leap year. If a year is divisible by 4 then it must also either not be
// divisible by 100 or if it is divisible by 100 it has to also be divisible by 400.
func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}

func generateRandomNumberInclusive(lowest int, highest int) int {
	var (
		difference          = highest - lowest
		offsetToBeInclusive = difference + 1
	)
	return rand.Intn(offsetToBeInclusive) + lowest
}

// Given a year, and a month (1-12) this function returns the count of days in that month. It uses
// the year to determine if February should get an extra day because it's a leap year.
func getDaysInMonth(year int, month int) int {
	switch month {
	case 2: // February
		if isLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

var era = "AD"

func main() {
	for count := 10; count > 0; count-- {
		year := generateRandomNumberInclusive(2025, 3025)
		month := generateRandomNumberInclusive(1, 12)
		daysInMonth := getDaysInMonth(year, month)
		day := generateRandomNumberInclusive(1, daysInMonth)
		fmt.Println(era, year, month, day)
	}
}
