package main

import "fmt"

func main() {
	var administrator *string
	fmt.Println(administrator)
	fmt.Println("Cannot dereference a nil pointer...")
	// fmt.Println(*administrator)

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(administrator)
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(administrator)
	fmt.Println(*administrator)

	bolden = bolden + " Jr."
	fmt.Println(administrator)
	fmt.Println(*administrator)

	*administrator = "Maj. Gen. " + bolden
	fmt.Println(administrator)
	fmt.Println(*administrator)

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(administrator)
	fmt.Println(*administrator)
	fmt.Println(administrator == major) // Prints true

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator)
	fmt.Println(*administrator)
	fmt.Println(administrator == major) // Prints false

	charles := *major         // A copy is made by assigning the dereferenced value
	*major = "Charles Bolden" // Replaces the string where the major pointer is pointing with "Charles Bolden"
	fmt.Println(charles)      // Prints Major General Charles Frank Bolden Jr. because that was the value at the time of copy
	fmt.Println(bolden)       // Prints Charles Bolden

	charles = "Charles Bolden"       // This string is identical to the value of "bolden"
	fmt.Println(charles == bolden)   // prints true
	fmt.Println(&charles == &bolden) // prints false
}
