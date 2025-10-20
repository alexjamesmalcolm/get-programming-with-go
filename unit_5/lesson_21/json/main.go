package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes)) // Because of the struct tags earlier we marshal to latitude and longitude!
}
