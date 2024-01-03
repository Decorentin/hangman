package hangmanpackage

import "fmt"

// function to print help menu
func HelpMenu(data *HangManData) {
	fmt.Println()
	fmt.Println("To play the hangman game, write the name of a .txt file as first argument. For example: ")
	fmt.Println("go run main/main.go words.txt")
	fmt.Println()
	fmt.Println("If you want to quit, write QUIT ")
	fmt.Println()
	fmt.Println("If you want to save, write STOP. To load the backup, write: ")
	fmt.Println("go run main/main.go --startWith save.txt")
	fmt.Println()
	fmt.Println("Enjoy!!")
	fmt.Println()
}
