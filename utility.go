package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Partner struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

func (app *App) reader() {
	filePath := os.Getenv("JSON_FILE_PATH")

	mode := os.Getenv("MODE")

	//ope file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when we're done with it

	// Create a decoder to read from the file
	decoder := json.NewDecoder(file)

	// Create a variable to store the decoded JSON data
	var data []Partner

	// Decode the JSON data into a slice of structs
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Now, you can work with the data as needed
	for _, item := range data {
		var domain string
		if mode == "mainnet" {
			switch item.Name {
			case "alchemy":
				domain = item.Domain + "/" + os.Getenv("ALCHEMY_KEY")
			default:
				domain = item.Domain
			}
		} else {
			domain = item.Domain
		}
		app.makeExternalCall(domain, item.Name)
	}
}
