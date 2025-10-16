package main

import "fmt"

func main() {
	peace := "shalom😀"
	for i := 0; i < len(peace); i++ {
		character := peace[i]
		fmt.Printf("%3v - %[1]c\n", character)
	}
}
