package main

import "fmt"

func cipher(message string) string {
	var characters []rune
	for _, char := range message {
		var (
			isLower = char >= 'a' && char <= 'z'
			isUpper = char >= 'A' && char <= 'Z'
		)
		if isLower || isUpper {
			char -= 13
			if (isLower && char < 'a') || (isUpper && char < 'A') {
				char += 26
			}
		}
		characters = append(characters, char)
	}
	return string(characters)
}

func decipher(message string) string {
	var characters []rune
	for _, char := range message {
		var (
			isLower = char >= 'a' && char <= 'z'
			isUpper = char >= 'A' && char <= 'Z'
		)
		if isLower || isUpper {
			char += 13
			if (isLower && char > 'z') || (isUpper && char > 'Z') {
				char -= 26
			}
		}
		characters = append(characters, char)
	}
	return string(characters)
}

func main() {
	message := "Hola Estación Espacial Internacional"
	fmt.Println(message)
	encodedMessage := cipher(message)
	fmt.Println(encodedMessage)
	decodedMessage := decipher(encodedMessage)
	fmt.Println(decodedMessage)
}
