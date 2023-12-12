package hangmanpackage

import (
	"fmt"
)

// function that manages part of the user interface
func UserInterface(data *HangManData) {
	if data.Attempts > 0 {
		fmt.Printf("Good Luck, you have, %d attempts. \n", data.Attempts)
	} else {
		data.Attempts = 10
		fmt.Printf("Good Luck, you have, %d attempts. \n", data.Attempts)
		fmt.Println(data.Word)
		fmt.Println()
	}
	GameBoucle(data)
}
