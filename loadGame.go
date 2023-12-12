package hangmanpackage

import (
	"encoding/json"
	"log"
	"os"
)

// function to load game data when the user wants to load a backup
func loadGame(filename string, data *HangManData) {
	datafile, err := os.ReadFile(filename) // Read the JSON data from the file
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(datafile, data) // Unmarshal the JSON data into the HangManData structure
	if err != nil {
		log.Fatal(err)
	}
}
