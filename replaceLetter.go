package hangmanpackage

// Replace the letter of the partially filled word with the chosen letter if it is in the word
func ReplaceLetter(word, currentWord, letter string) (string, bool) {
	result := ""
	inTheWord := false

	for i := 0; i < len(word); i++ {
		if string(word[i]) == letter {
			result += letter
			inTheWord = true
		} else {
			result += string(currentWord[i])
		}
	}

	if len(currentWord) > len(word) {
		result += currentWord[len(word):]
	}

	return result, inTheWord
}
