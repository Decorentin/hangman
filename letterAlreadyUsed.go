package hangmanpackage

// function that returns true if the letter entered by the user is present in the table of letters already proposed
func LetterAlreadyUsed(tabInput []string, input string) bool {
	for _, letter := range tabInput {
		if letter == input {
			return true
		}
	}
	return false
}
