package hangmanpackage

import (
	"fmt"
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
	for i := 0; i < numbreveal; i++ {
		max := len(words) - 1
		indexrandomword := rand.Intn(max)
		randomletter := string(words[indexrandomword])
		for j := 0; j < len(words); j++ {
			if string(words[j]) == randomletter {
				tabword[j] = randomletter
			}
		}
	}
	fmt.Println(strings.Join(tabword, ""))
	return strings.Join(tabword, "")
}
