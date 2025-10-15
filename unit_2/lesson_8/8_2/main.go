package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(86400)
	fmt.Println(first)

	second := new(big.Int)
	second.SetString("86400", 10)
	fmt.Println(second)
}
