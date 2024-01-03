package hangmanpackage

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// function to generate a random word in the file specified by the user
func RandomWord() string {
	dictionaryPath := os.Args[1]
	var tabword []string
	readFile, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile) // Crée un scanner pour lire le fichier.
	fileScanner.Split(bufio.ScanLines)        // Divise le fichier en lignes.
	// Parcours chaque ligne du fichier.
	for fileScanner.Scan() {
		// Ajoute le texte de la ligne actuelle à la slice trouverlemot.
		tabword = append(tabword, fileScanner.Text())
	}
	readFile.Close()
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator
	max := len(tabword)
	indexrandomword := rand.Intn(max)
	randomword := tabword[indexrandomword]

	return randomword
}
