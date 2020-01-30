package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	possibilities := [3]string{"rock", "paper", "scissors"}

	roshambo := choseRandom()
	guess := getGuess()

	fmt.Println("I chose ", possibilities[roshambo])

	guessIndex := findGuessIndex(guess, possibilities)
	userWins := didUserWin(roshambo, guessIndex)

	if guessIndex == roshambo {
		fmt.Println("Looks like we tied. Lame.")
	} else if isValidGuess(possibilities, guess) {
		if userWins {
			fmt.Println("I lost. Something must be wrong with my programming.")
		} else {
			fmt.Println("Loser User!!!")
			fmt.Printf("%s beats %s!", possibilities[roshambo], guess)
		}
	} else {
		fmt.Println("Your entry was invalid. CHEATER! I win by default!")
	}

}

func didUserWin(roshambo int, guessIndex int) bool {
	var userWon bool
	if roshambo == 0 && guessIndex == 2 {
		userWon = false
	} else if roshambo == 2 && guessIndex == 0 {
		userWon = true
	} else {
		userWon = roshambo < guessIndex
	}
	return userWon
}

func choseRandom() int {
	rand.Seed(time.Now().UnixNano())
	max := 99
	min := 1
	myRandom := rand.Intn((max - min + 1) + 0)
	roshambo := myRandom % 3
	return roshambo
}

func getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What's it gonna be? ")
	guess, _ := reader.ReadString('\n')
	return strings.TrimSpace(guess)
}

func isValidGuess(possibilities [3]string, guess string) bool {
	var isValid = false
	switch guess {
	case
		possibilities[0],
		possibilities[1],
		possibilities[2]:
		isValid = true
	}
	return isValid
}

func findGuessIndex(guess string, possibilities [3]string) int {
	var index int
	for i, possibility := range possibilities {
		if strings.Contains(guess, possibility) {
			index = i
		}
	}
	return index
}
