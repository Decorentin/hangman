package hangmanpackage

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// function to display the hangman according to the number of tries remaining
func HangmanPositions(data *HangManData) {
	f, err := os.Open("hangman.txt") //open "hangman.txt" because hangman positions are stocked in
	var tab []string
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tab = append(tab, scanner.Text()) //adds the contents of hangman.txt to the slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	if data.Attempts == 9 { //print the first position of hangman if Attempts is equal to 9
		for i := 0; i <= 7; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 8 { //print the seconde position of hangman if Attempts is equel to 8
		for i := 8; i <= 15; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 7 {
		for i := 16; i <= 23; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 6 {
		for i := 24; i <= 31; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 5 {
		for i := 32; i <= 39; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 4 {
		for i := 40; i <= 47; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 3 {
		for i := 48; i <= 55; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 2 {
		for i := 56; i <= 63; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 1 {
		for i := 64; i <= 71; i++ {
			fmt.Println(tab[i])
		}
	}
	if data.Attempts == 0 {
		for i := 72; i <= 79; i++ {
			fmt.Println(tab[i])
		}
	}
}
