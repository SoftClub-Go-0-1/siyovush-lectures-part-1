package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello! This is a simple guessing game")
	fmt.Println("Give me a number (> 0) and I will choose a number between 1 and your number (inclusive) and you try to guess it")
	fmt.Print("\nMax number I can choose: ")
	var maxNumber int
	fmt.Scan(&maxNumber)

	if maxNumber < 1 {
		fmt.Println("I won't play with you :( You are breaking the rules!")
		return
	}

	chosenNumber := rand.Intn(maxNumber) + 1

	fmt.Printf("Okay, I have chosen a number between 1 and %d. Try to guess it :)\n", maxNumber)
	var try, moves int

	for try != chosenNumber {
		fmt.Print("\nYour guess: ")
		fmt.Scan(&try)
		moves++
		if try > chosenNumber {
			fmt.Println("No, your number is too big. Try a smaller one.")
		} else if try < chosenNumber {
			fmt.Println("No, your number is too small. Try a bigger one.")
		} else {
			fmt.Printf("\nCorrect! It is %d. You won in %d moves.\n", chosenNumber, moves)
			break
		}
	}
}
