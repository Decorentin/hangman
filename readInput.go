package hangmanpackage

import "fmt"

func ReadInput() string {
	fmt.Print("Choose: ")
	var userInput string
	_, err := fmt.Scan(&userInput) // scan user input
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée :", err)
	}
	return userInput
}
