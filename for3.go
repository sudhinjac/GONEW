package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!!!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can You Guess what it is ? ")

	var guess int

	for {

		fmt.Println("Enter Your Guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congradulations! You guessed it right!")
			break
		} else if guess < target {
			fmt.Println("Too Low! Try Guessing a higer number")

		} else {
			fmt.Println("Too High! Try guessing a low number")
		}
	}
}
