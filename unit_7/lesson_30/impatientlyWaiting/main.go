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
	timeout := time.After(3 * time.Second)
	for range 5 {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}

}
