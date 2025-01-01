# Number Guessing Game

> Project Practice: https://roadmap.sh/projects/number-guessing-game

This is a simple number guessing game built in Go. The game generates a random number within a specified range, and the player needs to guess the number within a limited number of attempts.

## How to Play

1. **Choose Difficulty:** Select a difficulty level (Easy, Medium, or Hard). This determines the number of chances you have to guess the number.
2. **Start Guessing:** Enter your guess when prompted.
3. **Feedback:** The game will provide feedback on whether your guess is too high or too low.
4. **Win or Lose:** If you guess the correct number within the allowed attempts, you win! Otherwise, you lose.

## Features

- **Three Difficulty Levels:** Easy, Medium, and Hard.
- **Random Number Generation:** The game generates a random secret number for each round.
- **User Input Validation:** Ensures valid numerical input from the player.
- **Clear Feedback:** Provides clear feedback on guess accuracy.
- **Game Over Message:** Displays a game over message when the player runs out of chances.

## How to Run

1. Make sure you have Go installed on your system.
2. Clone this repository: `git clone <repository-url>`
3. Navigate to the project directory: `cd <project-directory>`
4. Build the game: `go build`
5. Run the game: `./<executable-name>`

## Example Gameplay

```
Welcome to the Number Guessing Game!
Select a difficulty level:

Easy (10 chances)
-> Medium (5 chances) <-
Hard (3 chances)

Great! You have selected the Medium difficulty level. Let's start the game!

Enter your guess: 50
Incorrect! The number is greater than 50. 

Enter your guess: 75 
Incorrect! The number is less than 75.

Enter your guess: 60 
Incorrect! The number is greater than 60. 

Enter your guess: 65 
Incorrect! The number is greater than 65. 

Enter your guess: 70 
Incorrect! The number is greater than 70.

Game Over! You ran out of chances.
```