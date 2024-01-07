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
		fmt.Println(actualword)
		return actualword
	}
	fmt.Println(actualword)
	return actualword
}
