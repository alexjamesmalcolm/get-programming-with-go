package main

import (
	"fmt"
	"math/rand"
	"time"
)

func countdown() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count -= 1
	}
	fmt.Println("Liftoff!")
}

func infinity() {
	var degrees = 0
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

func generateRandomNumberInclusive(lowest int, highest int) int {
	var (
		difference          = highest - lowest
		offsetToBeInclusive = difference + 1
	)
	return rand.Intn(offsetToBeInclusive) + lowest
}

func countdownWithFailureChance() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			fmt.Println("Abort!")
			return
		}
		count -= 1
	}
	fmt.Println("Liftoff!")
}

func guessingGame(answer int) {
	var (
		maximum = 100
		minimum = 1
	)
	if answer > maximum || answer < minimum {
		fmt.Println(
			"The number you supplied was",
			answer, "but this program only tries to guess answers ",
			minimum, "-", maximum,
		)
		fmt.Println("Please enter another number within that range.")
	}
	var guess = generateRandomNumberInclusive(minimum, maximum)
	for guess != answer {
		if guess > answer {
			fmt.Println("Guessed", guess, "but that was too high!")
			maximum = guess - 1
		} else {
			fmt.Println("Guessed", guess, "but that was too low!")
			minimum = guess + 1
		}
		guess = generateRandomNumberInclusive(minimum, maximum)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Your number is:", answer)
}

func main() {
	var answer = generateRandomNumberInclusive(1, 100)
	fmt.Println("The secret number that I'm thinking of is", answer)
	guessingGame(answer)
}
