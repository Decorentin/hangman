package main

import (
	"fmt"
	"github.com/Decorentin/hangmanpackage"
)

func main() {
	data := hangmanpackage.HangManData{}

	data.ToFind = hangmanpackage.RandomWord()
	data.Word = hangmanpackage.RandomLetters(data.ToFind)
	data.Attempts = 10

	for data.ToFind != data.Word && data.Attempts > 0 {
		data.Input = hangmanpackage.ReadInput()
		newWord, newAttempts := hangmanpackage.ModifyWord(data.ToFind, data.Word, data.Input, data.Attempts)
		data.Word = newWord
		data.Attempts = newAttempts
	}

	if data.ToFind == data.Word {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}

