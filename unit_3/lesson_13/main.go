package main

import "fmt"

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

func main() {
	var kelvin kelvin = 294.0
	celsius := kelvin.celsius()
	fmt.Printf("%.2f° K is %.2f° C\n", kelvin, celsius)
	fahrenheit := celsius.fahrenheit()
	fmt.Printf("%.2f° C is %.2f° F\n", celsius, fahrenheit)
	fahrenheit = kelvin.fahrenheit()
	fmt.Printf("%.2f° K is %.2f° F\n", kelvin, fahrenheit)

	kelvin = 233
	celsius = kelvin.celsius()
	fmt.Printf("%.2f° K is %.2f° C\n", kelvin, celsius)
	fahrenheit = celsius.fahrenheit()
	fmt.Printf("%.2f° C is %.2f° F\n", celsius, fahrenheit)
	fahrenheit = kelvin.fahrenheit()
	fmt.Printf("%.2f° K is %.2f° F\n", kelvin, fahrenheit)
}
