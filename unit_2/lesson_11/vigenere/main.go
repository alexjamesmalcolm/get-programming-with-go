package main

import (
	"errors"
	"fmt"
	"strings"
)

func valueOfRune(r rune) (int32, error) {
	if r >= 'A' && r <= 'Z' {
		return int32(r - 'A'), nil
	}
	return 0, errors.New("Cannot get value of " + string(r))
}

func vigenereDecipher(cipher string, keyword string) (string, error) {
	keywordRunes := []rune(keyword)
	var decoded []rune
	for i, char := range cipher {
		keywordIndex := i % len(keywordRunes)
		keywordChar := keywordRunes[keywordIndex]
		valueOfKeywordChar, err := valueOfRune(keywordChar)
		if err != nil {
			return "", err
		}
		shiftedChar := char - valueOfKeywordChar
		if shiftedChar < 'A' {
			shiftedChar += 26
		}
		decoded = append(decoded, shiftedChar)
	}
	return string(decoded), nil
}

func vigenereCipher(message string, keyword string) (string, error) {
	message = strings.ReplaceAll(strings.ToUpper(message), " ", "")
	keywordRunes := []rune(keyword)
	var cipher []rune
	for i, char := range message {
		keywordIndex := i % len(keywordRunes)
		keywordChar := keywordRunes[keywordIndex]
		valueOfKeywordChar, err := valueOfRune(keywordChar)
		if err != nil {
			return "", err
		}
		shiftedChar := char + valueOfKeywordChar
		if shiftedChar > 'Z' {
			shiftedChar -= 26
		}
		cipher = append(cipher, shiftedChar)
	}
	return string(cipher), nil
}

func main() {
	fmt.Println(valueOfRune('A'))
	fmt.Println(valueOfRune('Z'))
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	decoded, err := vigenereDecipher(cipherText, keyword)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cipherText, "decoded using keyword", keyword, "is:", decoded)

	plainText := "your message goes here"
	cipher, err := vigenereCipher(plainText, keyword)
	fmt.Println(plainText, "encoded using keyword", keyword, "is:", cipher)
	decoded, err = vigenereDecipher(cipher, keyword)
	fmt.Println(cipher, "decoded using keyword", keyword, "is:", decoded)
}
