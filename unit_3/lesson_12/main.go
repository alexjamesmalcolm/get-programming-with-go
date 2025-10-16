package main

import "fmt"

func kelvinToCelsius(kelvin float64) (celsius float64) {
	return kelvin - 273.15
}

func celsiusToFahrenheit(celsius float64) (fahrenheit float64) {
	return (celsius * 9 / 5) + 32
}

func kelvinToFahrenheit(kelvin float64) (fahrenheit float64) {
	celsius := kelvinToCelsius(kelvin)
	return celsiusToFahrenheit(celsius)
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%.2f° K is %.2f° C\n", kelvin, celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f° C is %.2f° F\n", celsius, fahrenheit)
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Printf("%.2f° K is %.2f° F\n", kelvin, fahrenheit)

	kelvin = 233
	celsius = kelvinToCelsius(kelvin)
	fmt.Printf("%.2f° K is %.2f° C\n", kelvin, celsius)
	fahrenheit = celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f° C is %.2f° F\n", celsius, fahrenheit)
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Printf("%.2f° K is %.2f° F\n", kelvin, fahrenheit)
}
