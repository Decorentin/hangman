package hangmanpackage

import (
	"math/rand"
)

// function that calculates the number of letters to be revealed and prints the word with the revealed letters and "_".
func RandomLetters(data *HangManData) {
	nRandomLetters := (len(data.ToFind) / 2) - 1
	var tabLetters []string
	var word string
	for i := 0; i < len(data.ToFind); i++ {
		word += "_"
	}
	i := 0
	for i < nRandomLetters {
		randomIndex := rand.Intn(len(data.ToFind)) //choose random letters in the word
		randomLetter := string(data.ToFind[randomIndex])
		data.TabInput = append(data.TabInput, randomLetter)
		found := 0
		for _, letter := range tabLetters {
			if letter == randomLetter {
				found = 1 //Set 'found' to 1 if the letter is found in 'tabLetters
			}
		}
		if found != 1 {
			tabLetters = append(tabLetters, randomLetter) // replace the corresponding "_" in 'word',
			word = ReplaceLetter(data.ToFind, word, randomLetter)
			i++
		}
	}
	data.Word = word
}
