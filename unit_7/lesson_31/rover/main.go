package main

import (
	"fmt"
	"image"
	"time"
)

func worker() {
	pos := image.Point{10, 10}
	direction := image.Point{1, 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is", pos)
			next = time.After(time.Second)
		}
	}
}
func main() {
	go worker()
	time.Sleep(10 * time.Second)
}
