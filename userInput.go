package hangmanpackage

import (
	"fmt"
	"strings"
)

// function that manages user input
func UserInput(data *HangManData) {
	fmt.Print("Choose: ")
	_, err := fmt.Scan(&data.Input) //scan user input
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée :", err)
		return
	}
	if data.Input == "STOP" { //if input is STOP, game is save in "save.txt"
		saveGame("save.txt", *data)
		fmt.Println("Game Saved in save.txt.")
		data.StopGame = true
		return
	}
	if data.Input == "QUIT" { //if input is QUIT, game is over
		fmt.Println("You've decided to stop playing")
		data.StopGame = true
		return
	}
	if len(data.Input) > 1 { //if input is > 1, func SuggestWord
		SuggestWord(data)
	} else {
		data.OneLetter = true
	}
	if len(data.Input) < 1 { //if input < 1, error message
		fmt.Println("Please enter a letter or word")
	}
	checkletter := strings.Contains(data.ToFind, data.Input) //check if letter input is in the word
	if !data.Won && data.OneLetter {
		if LetterAlreadyUsed(data) { //checks if the letter input has already been used
			fmt.Println("Error: you have already proposed this letter")
			fmt.Println()
			fmt.Println(data.Word)
			fmt.Println()
			return
		}
		data.TabInput = append(data.TabInput, data.Input) //adds the letter input to the sutilised letter table
		if checkletter {
			result := ""
			for i := 0; i < len(data.ToFind); i++ { //adds the letter input to the word
				if string(data.ToFind[i]) == data.Input {
					result += data.Input
				} else {
					result += string(data.Word[i])
				}
			}
			data.Word = result
		} else {
			data.Attempts--
			fmt.Printf("Not present in the word, %d attempts remaining\n", data.Attempts) //Attempts - 1 if the letter input is not in the word
			HangmanPositions(data)
		}
	}
	if checkletter && !data.Won && data.OneLetter { //print the word (the actual word not the one to find) if the game is not won
		fmt.Println(data.Word)
		fmt.Println()
	}
}
