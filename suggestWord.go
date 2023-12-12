package hangmanpackage

import "fmt"

// function allowing the user to propose a word, if the proposed word is the same as the one to be found, the user has won.
func SuggestWord(data *HangManData) {
	data.OneLetter = false
	if data.Input == data.ToFind { //if the word proposed by the player is the word to guess, its win
		data.Won = true
		fmt.Println()
		fmt.Println("Congrats !")
	} else { //if its not the word to guess, Attempts -1 if Attempts is not = 1 (to prevent it from being negative)
		if data.Attempts == 1 {
			data.Attempts--
			fmt.Printf("This is not the word, %d attempts remaining\n", data.Attempts)
			HangmanPositions(data) //print hangman
		} else {
			data.Attempts -= 2
			fmt.Printf("This is not the word, %d attempts remaining\n", data.Attempts) //else, Attempt -2
			HangmanPositions(data)
		}
	}
}
