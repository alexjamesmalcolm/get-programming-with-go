package main

import "fmt"

func decodeCaesar(encodedMessage string) string {
	var decodedCharacters []byte
	for i := range encodedMessage {
		char := encodedMessage[i]
		var decodedChar byte
		switch char {
		case ' ', ',', '.':
			decodedChar = char
		default:
			decodedChar = char - 3
		}
		decodedCharacters = append(decodedCharacters, decodedChar)
		fmt.Printf("%c\n", char)
	}
	return string(decodedCharacters)
}

func main() {
	const cipher = "L fdph, L vdz, L frqtxhuhg."
	decodedCipher := decodeCaesar(cipher)
	fmt.Println(decodedCipher, "—Julius Caesar")
}
