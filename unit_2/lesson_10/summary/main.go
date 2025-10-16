package main

import (
	"errors"
	"fmt"
)

func parseBoolean(value string) (bool, error) {
	switch value {
	case "true", "yes", "1":
		return true, nil
	case "false", "no", "0":
		return false, nil
	default:
		return false, errors.New("Value (" + value + ") cannot be converted to a boolean.")
	}
}

func main() {
	fmt.Println(parseBoolean("true"))
	fmt.Println(parseBoolean("yes"))
	fmt.Println(parseBoolean("1"))
	fmt.Println(parseBoolean("false"))
	fmt.Println(parseBoolean("no"))
	fmt.Println(parseBoolean("0"))
	fmt.Println(parseBoolean("sure"))
}
