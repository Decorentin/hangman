package hangmanpackage

import (
	"fmt"
	"strings"
)

// function that manages user input
func UserInput(input string, toFind string, word string, attempts int, tabInput []string, won bool, stopGame bool) (string, int, []string, bool, bool) {
	fmt.Print("Choose: ")
	_, err := fmt.Scan(&input) // scan user input
	if err != nil {
		fmt.Println("Error reading input:", err)
		return word, attempts, tabInput, won, stopGame
	}

	var oneLetter bool
	if len(input) > 1 { // if input is > 1, func SuggestWord
		attempts, oneLetter, won = SuggestWord(input, toFind, attempts, oneLetter, won)
	} else {
		oneLetter = true
	}

	checkletter := strings.Contains(toFind, input) // check if the letter input is in the word

	if !won && oneLetter {
		if LetterAlreadyUsed(tabInput, input) { // checks if the letter input has already been used
			fmt.Println("Error: you have already proposed this letter")
			fmt.Println()
			fmt.Println(word)
			fmt.Println()
			return word, attempts, tabInput, won, stopGame
		}

		tabInput = append(tabInput, input) // adds the letter input to the used letter table

		if checkletter {
			result := ""
			for i := 0; i < len(toFind); i++ { // adds the letter input to the word
				if string(toFind[i]) == input {
					result += input
				} else {
					result += string(word[i])
				}
			}
			word = result
		} else {
			attempts--
			fmt.Printf("Not present in the word, %d attempts remaining\n", attempts) // Attempts - 1 if the letter input is not in the word
		}
	}

	if checkletter && !won && oneLetter { // print the word (the actual word, not the one to find) if the game is not won
		fmt.Println(word)
		fmt.Println()
	}

	return word, attempts, tabInput, won, stopGame
}
