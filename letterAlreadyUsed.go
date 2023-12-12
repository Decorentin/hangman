package hangmanpackage

// function that returns true if the letter entered by the user is present in the table of letters already proposed
func LetterAlreadyUsed(data *HangManData) bool {
	for _, letter := range data.TabInput {
		if letter == data.Input {
			return true
		}
	}
	return false
}
