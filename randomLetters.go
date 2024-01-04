package hangmanpackage

import (
	"math/rand"
	"strings"
)

// Replace a random set of letters in the word with "_"
func RandomLetters(word string) string {
	numToReveal := (len(word) / 2) - 1
	revealedIndices := make(map[int]bool)

	// Generate unique random indices to reveal
	for len(revealedIndices) < numToReveal {
		index := rand.Intn(len(word))
		revealedIndices[index] = true
	}

	// Create the partially revealed word
	var result strings.Builder
	for i, char := range word {
		if revealedIndices[i] {
			result.WriteRune(char)
		} else {
			result.WriteRune('_')
		}
	}

	return result.String()
}
