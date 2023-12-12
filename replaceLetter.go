package hangmanpackage

// Replace the letter of the partially filled word with the chosen letter if it is in the word
func ReplaceLetter(word, currentWord, letter string) string {
	result := ""
	for i := 0; i < len(word); i++ {
		if string(word[i]) == letter {
			result += letter
		} else {
			result += string(currentWord[i])
		}
	}
	return result
}
