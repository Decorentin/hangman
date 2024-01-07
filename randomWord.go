package hangmanpackage

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// function to generate a random word in the file specified by the user
func RandomWord() string {
	file := os.Args[1]
	content, err := os.ReadFile(file) //opens user-selected txt
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(content), "\n") //choose a random word in the txt
	randomWord := words[rand.Intn(len(words))]
	fmt.Println(randomWord)
	return randomWord
}
