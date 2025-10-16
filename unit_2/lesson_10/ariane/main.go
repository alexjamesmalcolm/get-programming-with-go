package main

import (
	"errors"
	"fmt"
	"math"
)

func convertFloat64ToInt16(value float64) (int16, error) {
	if value < math.MinInt16 {
		return 0, errors.New("The float64 is too small to be contained inside of an int16.")
	} else if value > math.MaxInt16 {
		return 0, errors.New("The float64 is too big to be contained inside of an int16.")
	}
	return int16(value), nil
}

func main() {
	var bh float64 = 32767 + 1
	var h, error = convertFloat64ToInt16(bh)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(h)
}
