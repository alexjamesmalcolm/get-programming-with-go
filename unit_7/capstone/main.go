package main

import (
	"capstone/location"
	"capstone/mars"
	"capstone/radio"
	"capstone/rover"
	"fmt"
	"image"
	"time"
)

func main() {
	grid := mars.NewMarsGrid(100, 100)
	earth := radio.NewRadio("Earth", &radio.Antenna{
		SignalRange: 500_000_000,
		Distancer: location.Location{
			Planet:    location.Earth,
			OnSurface: true,
			Coordinate: location.Coordinate{
				Lat:  35.4259,
				Long: -116.8889,
			},
		},
		DataRate: 0,
	})
	_, err := rover.NewRover(
		"Perseverance",
		grid,
		image.Point{25, 25},
		location.Location{
			Coordinate: location.Coordinate{
				Lat:  18.447,
				Long: 77.402,
			},
			Planet:    location.Mars,
			OnSurface: true,
		},
	)
	if err != nil {
		fmt.Println("There was an error initializing Perseverance")
		return
	}
	rover.NewRover(
		"Curiosity",
		grid,
		image.Point{75, 75},
		location.Location{
			Coordinate: location.Coordinate{
				Lat:  -4.80783741,
				Long: 137.38103,
			},
			Planet:    location.Mars,
			OnSurface: true,
		},
	)
	radio.NewRadio("Maven", &radio.Antenna{
		DataRate:    6e+6,
		SignalRange: 500_000_000,
		Distancer: location.Location{
			Planet:    location.Mars,
			OnSurface: false,
		},
	})
	radio.NewRadio("MRO", &radio.Antenna{
		DataRate:    6e+6,
		SignalRange: 500_000_000,
		Distancer: location.Location{
			Planet:    location.Mars,
			OnSurface: false,
		},
	})

	fmt.Println(earth)
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		messages := earth.Inbox()
		if len(messages) > 0 {
			// fmt.Printf("Earth has received messages!\n", messages)
			fmt.Println("Earth has received messages!")
			for _, m := range messages {
				fmt.Println(m)
			}
		}
	}
}
