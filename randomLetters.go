package hangmanpackage

import (
	"math/rand"
	"strings"
)

// function that calculates the number of letters to be revealed and prints the word with the revealed letters and "_".
func RandomLetters(words string) string {
	numbreveal := (len(words) / 2) - 1
	tabword := make([]string, len(words))
	for i := range tabword {
		tabword[i] = "_"
	}

	// Utilise un ensemble pour suivre les lettres déjà révélées
	revealedLetters := make(map[string]bool)

	for i := 0; i < numbreveal; i++ {
		// Choisis une lettre aléatoire qui n'a pas encore été révélée
		randomIndex := rand.Intn(len(words))
		for revealedLetters[string(words[randomIndex])] {
			randomIndex = rand.Intn(len(words))
		}

		randomLetter := string(words[randomIndex])
		revealedLetters[randomLetter] = true

		// Remplace la première occurrence non révélée dans tabword
		for j := 0; j < len(words); j++ {
			if string(words[j]) == randomLetter && tabword[j] == "_" {
				tabword[j] = randomLetter
				break
			}
		}
	}

	return strings.Join(tabword, "")
}
