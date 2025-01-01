package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func printWelcomeMsg() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("")
}

func selectDifficultyLevel() DifficultyLevel {
	formattedLevels := []string{}
	for i, lvl := range difficultyLevels {
		formattedLevels = append(formattedLevels, fmt.Sprintf("%d. %s (%d chances)", i, lvl.Name, lvl.Chances))
	}

	prompt := promptui.Select{
		Label: "Please select the difficulty level",
		Items: formattedLevels,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return difficultyLevels[index]
}
