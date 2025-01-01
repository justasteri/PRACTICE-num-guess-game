package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DifficultyLevel struct {
	Name    string
	Chances int
}

var difficultyLevels = []DifficultyLevel{
	{Name: "Easy", Chances: 10},
	{Name: "Medium", Chances: 5},
	{Name: "Hard", Chances: 3},
}

func main() {
	printWelcomeMsg()

	randomNumber := genRandNum(1, 100)

	selectedDifLvl := selectDifficultyLevel()

	fmt.Printf("Great! You have selected the %s difficulty level.\n", selectedDifLvl.Name)
	fmt.Println("Let's start the game!")

	startGame(randomNumber, selectedDifLvl)
}

func genRandNum(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min+1) + min
}
