package main

import (
	"fmt"
	"sync"
	"time"
)

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

// NewVisited instantiates a new Visited
func NewVisited() Visited {
	return Visited{visited: make(map[string]int)}
}

// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	googleURL := "https://google.com"
	v := NewVisited()
	go v.VisitLink(googleURL)
	go v.VisitLink(googleURL)
	time.Sleep(200 * time.Millisecond)
	fmt.Println(v.visited)
}
