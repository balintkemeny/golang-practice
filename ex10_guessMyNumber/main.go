package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var inRange int
	fmt.Println("Please specify range:")
	fmt.Scanf("%d", &inRange)

	var answer int = rand.Intn(inRange)
	var lives int = 5

	for {
		fmt.Printf("Guess a number! You have %d lives left.\n", lives)
		var guess int
		fmt.Scanf("%d", &guess)

		if guess == answer {
			fmt.Println("Congratulations, you guessed the correct number")
			break
		} else if guess < answer {
			fmt.Println("Your guess is too low.")
			lives--
		} else {
			fmt.Println("Your guess is too high.")
			lives--
		}

		if lives < 1 {
			fmt.Println("You've lost the game. Better luck next time!")
			break
		}
	}
}
