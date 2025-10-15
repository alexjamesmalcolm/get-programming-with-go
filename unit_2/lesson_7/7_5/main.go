package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red) // Prints 0
	var number int8 = 127
	number++
	fmt.Println(number) // Prints -128
	var bigNumber uint16 = 65535
	bigNumber++
	fmt.Println(bigNumber) // Prints 0
	var zero uint8 = 0
	zero--
	fmt.Println(zero) // Prints 255
}
