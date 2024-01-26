package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type RequestPayload struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// custom json response type

func (app *App) makeExternalCall(url, partner string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the wait group counter when the function completes

	requestBody := RequestPayload{
		Jsonrpc: "2.0",
		Method:  "web3_clientVersion",
		Params:  []interface{}{}, // An empty array as Params
		ID:      1,
	}

	// Create some JSON to send to the external service
	jsonData, _ := json.MarshalIndent(requestBody, "", "\t")

	// Create a request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Create a client and make the request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read and log the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("--------------------------------------------------")
	log.Printf("%-50s %s", "Partner:", partner)
	log.Printf("%-50s %s", "Network:", os.Getenv("MODE"))
	log.Printf("%-50s %s", "Domain:", url)
	log.Printf("%-50s %s", "Response Body:", string(responseBody))
	log.Println("--------------------------------------------------")
}
