package hangmanpackage

import "fmt"

// loop that manages the game scenario, the game stops if the word has been found or if the number of tries has reached 0
func GameBoucle(data *HangManData) {
	if data.GameRestart {
		fmt.Printf("Welcome Back, you have , %d attempts remaining\n", data.Attempts) //only if the game restart after get load
		fmt.Println(data.Word)
		fmt.Println()
	}
	for data.ToFind != data.Word || data.Attempts > 0 { //loop that continues until the word has been or if the counter is equal to zero
		if data.Won || data.StopGame {
			break
		}
		UserInput(data)
		if data.ToFind == data.Word { //word guess
			fmt.Println("Congrats !")
			break
		}
		if data.Attempts <= 0 { //no more attempts
			fmt.Println("Loose")
			fmt.Println("The word was:", data.ToFind)
			break
		}
	}
}
