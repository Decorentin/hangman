package hangmanpackage

import "fmt"

// loop that manages the game scenario, the game stops if the word has been found or if the number of tries has reached 0
// loop that manages the game scenario, the game stops if the word has been found or if the number of tries has reached 0
func GameBoucle(data *HangManData) {
	if data.GameRestart {
		fmt.Printf("Welcome Back, you have %d attempts remaining\n", data.Attempts) // seulement si le jeu redémarre après avoir été chargé
		fmt.Println(data.Word)
		fmt.Println()
	}

	// Utilisez une boucle infinie pour gérer le jeu jusqu'à ce qu'il soit terminé
	for {
		// Break si le joueur a gagné ou a choisi d'arrêter le jeu
		if data.Won || data.StopGame {
			break
		}

		// Utilisez la fonction UserInput avec les données actuelles du jeu
		data.Word, data.Attempts, data.TabInput, data.Won, data.StopGame = UserInput(data.Input, data.ToFind, data.Word, data.Attempts, data.TabInput, data.Won, data.StopGame)

		// Conditions pour gérer différentes situations du jeu
		if data.ToFind == data.Word {
			fmt.Println("Congrats !")
			break
		}

		if data.Attempts <= 0 {
			fmt.Println("Loose")
			fmt.Println("The word was:", data.ToFind)
			break
		}
	}
}
