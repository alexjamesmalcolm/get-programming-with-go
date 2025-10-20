package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate returns a fictional measure of time for a given date.
func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24
	return 1000 + day + h
}

type sol int

func (s sol) YearDay() int { return int(s % 668) }
func (s sol) Hour() int    { return 0 }

func main() {
	day := time.Date(2025, 10, 20, 7+12, 16, 0, 0, time.Local)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))
}
