package main

import (
	"fmt"
	"strings"
)

type kelvin float64
type celsius float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	c := k.celsius()
	return c.fahrenheit()
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9 / 5) + 32)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

type getRowFn = func(row int) (float64, float64)

func drawRow(a, b float64) {
	fmt.Printf("| %5.1f\t| %5.1f\t|\n", a, b)
}

func drawTable(headerA, headerB string, rowCount int, getRow getRowFn) {
	lineBreak := strings.Repeat("=", 17)
	fmt.Println(lineBreak)
	fmt.Printf("| %v\t| %v\t|\n", headerA, headerB)
	fmt.Println(lineBreak)
	for i := range rowCount {
		drawRow(getRow(i))
	}
	fmt.Println(lineBreak)
}

func main() {
	var (
		getCelsiusToFahrenheitRow getRowFn = func(row int) (float64, float64) {
			c := celsius(-40 + (row * 5))
			return float64(c), float64(c.fahrenheit())
		}
		getFahrenheitToCelsiusRow getRowFn = func(row int) (float64, float64) {
			f := fahrenheit(-40 + (row * 5))
			return float64(f), float64(f.celsius())
		}
	)
	rowCount := 29
	drawTable("°C", "°F", rowCount, getCelsiusToFahrenheitRow)
	fmt.Println()
	drawTable("°F", "°C", rowCount, getFahrenheitToCelsiusRow)
}
