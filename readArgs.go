package hangmanpackage

import (
	"log"
	"os"
)

// function that analyzes the command(s) and executes functions based on these arguments
func ReadArgs(data *HangManData) {
	if len(os.Args) < 2 || len(os.Args) > 3 || os.Args[1] == "--help" { //returns to the help menu if there's an argument problem
		HelpMenu(data)
		return
	}
	if os.Args[1] == "--startWith" { //if the second argument is "--startWith", the game will be started with a save game
		if len(os.Args) < 3 {
			log.Fatal("Usage: ./hangman --startWith <filename>")
		}
		filename := os.Args[2]
		loadGame(filename, data)
		data.GameRestart = true
		GameBoucle(data)
	} else { //else the game launches normally
		RandomWord(data)
		RandomLetters(data)
		UserInterface(data)
	}
}
