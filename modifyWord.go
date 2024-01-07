package hangmanpackage

import (
	"fmt"
	"strings"
)

func ModifyWord(wordtofind, actualword, letter string) string {
	// Vérifie si la lettre est présente dans le mot
	if strings.Contains(wordtofind, letter) {
		// Met à jour actualword en ajoutant la lettre à la position correspondante
		for i, r := range wordtofind {
			if r == rune(letter[0]) {
				actualword = actualword[:i] + letter + actualword[i+1:]
			}
		}

		fmt.Println("Dans le mot")
		return actualword
	}

	// Si la lettre n'est pas dans le mot, retourne le mot inchangé
	fmt.Println("Pas dans le mot")
	return actualword
}
