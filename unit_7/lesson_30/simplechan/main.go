package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	c <- id
}

func main() {
	c := make(chan int)
	for i := range 5 {
		go sleepyGopher(i, c)
	}
	for range 5 {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}
