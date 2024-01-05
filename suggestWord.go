package hangmanpackage

import "fmt"

// SuggestWord allows the user to propose a word, if the proposed word is the same as the one to be found, the user has won.
func SuggestWord(input string, toFind string, attempts int, oneLetter bool, won bool) (newAttempts int, newOneLetter bool, newWon bool) {
	newOneLetter = false

	if input == toFind { // if the word proposed by the player is the word to guess, they win
		newWon = true
		fmt.Println()
		fmt.Println("Congrats!")
	} else { // if it's not the word to guess, Attempts -1 if Attempts is not = 1 (to prevent it from being negative)
		if attempts == 1 {
			newAttempts = attempts - 1
			fmt.Printf("This is not the word, %d attempts remaining\n", newAttempts)
			// Print hangman positions
		} else {
			newAttempts = attempts - 2
			fmt.Printf("This is not the word, %d attempts remaining\n", newAttempts)
			// Print hangman positions
		}
	}

	return newAttempts, newOneLetter, newWon
}
