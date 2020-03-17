package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number")

	//random number generator
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(50) //generates numbers between 0 and n (10)

	var guess int

	for{
		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too Big")
		} else if guess < secretNumber {
			fmt.Println("Too small")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}