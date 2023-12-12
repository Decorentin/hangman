package hangmanpackage

// structure named "HanManData" which allows variables to be used in various functions
type HangManData struct {
	Word        string   // Word composed of '_', ex: H_ll_
	ToFind      string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts    int      // Number of attempts left
	Input       string   //input enter by the user
	TabInput    []string // table that stores the letters already proposed
	Won         bool     //variable that becomes true when the word is guessed
	OneLetter   bool     //variable that becomes true when the input is one letter
	StopGame    bool     //variable for stopping the game
	GameRestart bool     //variable that becomes true when the game is launched from a savegame
}
