package main

import (
	"encoding/json"
	"fmt"
)

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	Degrees    float64 `json:"degrees"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Hemisphere byte    `json:"hemisphere"`
}

// decimal converts a d/m/s coordinate to a decimal degrees.
func (coord coordinate) decimal() float64 {
	sign := 1.0
	switch coord.Hemisphere {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (coord.Degrees + coord.Minutes/60 + coord.Seconds/3600)
}

func (c coordinate) String() string {
	return fmt.Sprintf(`%v°%v'%0.1f" %c`, c.Degrees, c.Minutes, c.Seconds, c.Hemisphere)
}

func (coord coordinate) MarshalJSON() ([]byte, error) {
	type data struct {
		Decimal    float64 `json:"decimal"`
		Dms        string  `json:"dms"`
		Degrees    float64 `json:"degrees"`
		Minutes    float64 `json:"minutes"`
		Seconds    float64 `json:"seconds"`
		Hemisphere string  `json:"hemisphere"`
	}
	return json.Marshal(data{
		Decimal:    coord.decimal(),
		Dms:        coord.String(),
		Degrees:    coord.Degrees,
		Minutes:    coord.Minutes,
		Seconds:    coord.Seconds,
		Hemisphere: string(coord.Hemisphere),
	})
}

func main() {
	c := coordinate{135, 54, 0, 'E'}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
