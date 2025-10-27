package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
}

func main() {
	for i := range 5 {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
