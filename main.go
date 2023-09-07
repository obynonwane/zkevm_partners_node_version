package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type App struct{}

func main() {
	app := App{}
	app.loadEnv()

	log.Println(os.Getenv("MODE"))
	app.reader()
}

func (app *App) loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
