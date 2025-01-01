package main

import "fmt"

func startGame(randomNumber int, difficulty DifficultyLevel) {
	msgIncorrectGuess := "Incorrect! The number is "
	var userGuess int

	for attempts := 1; attempts <= difficulty.Chances; attempts++ {
		fmt.Print("\nEnter your guess: ")
		_, err := fmt.Scan(&userGuess)
		if err != nil {
			fmt.Println("Error: Invalid number.")
			return
		}

		if userGuess == randomNumber {
			fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}

		stringComparator := "less"

		if randomNumber > userGuess {
			stringComparator = "greater"
		}

		fmt.Printf("%s %s %d.", msgIncorrectGuess, stringComparator, userGuess)
	}

	fmt.Println("\n\nGame Over! You ran out of chances.")
}
