package main

import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}

func main() {
	go sleepyGopher()           // go keyword is used to start a goroutine
	time.Sleep(4 * time.Second) // Needed if we want to wait for the sleepyGopher to return
}
