package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
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
			case "infura":
				domain = item.Domain + "/" + os.Getenv("INFURA_API_KEY")
			default:
				domain = item.Domain
			}
		} else {
			domain = item.Domain
		}

		var wg sync.WaitGroup
		wg.Add(1) //incrememts the wait group counter by 1 everytime

		// Run the makeExternalCall function as a goroutine
		go app.makeExternalCall(domain, item.Name, &wg)

		// Wait for the goroutine to complete i.e the counter to reach zero
		//blocks until the counter becomes zero
		wg.Wait()

	}
}
