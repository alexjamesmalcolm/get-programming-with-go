// Write a program that uses a loop to continuously append an element to a slice. Print out the
// capacity of the slice whenever it changes. Does append always double the capacity when the
// underlying array runs out of room?
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	const limit = 2000000000
	var (
		data     = []uint8{0}
		capacity = cap(data)
	)
	for len(data) < limit {
		data = append(data, uint8(rand.IntN(256)))
		if cap(data) != capacity {
			lastCap := capacity
			capacity = cap(data)
			length := len(data)
			fmt.Printf(
				"When data was %v long its capacity rose to %v. It grew by %v and multiplied by %.2f.\n",
				length,
				capacity,
				capacity-lastCap,
				float64(capacity)/float64(lastCap),
			)
			// The capacity of a slice that has a single value appended to it eventually increases
			// at a rate of 1.25x.
		}
	}
}
