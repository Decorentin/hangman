package hangmanpackage

import (
	"fmt"
	"strings"
)

func ModifyWord(wordtofind, actualword, letter string, attempts *int) string {
	if strings.Contains(wordtofind, letter) {
		for i, r := range wordtofind {
			if r == rune(letter[0]) {
				actualword = actualword[:i] + letter + actualword[i+1:]
			}
		}
		fmt.Println(*attempts)
		fmt.Println(actualword)
		return actualword
	}
	*attempts -= 1
	fmt.Println(*attempts)
	fmt.Println(actualword)
	return actualword
}
