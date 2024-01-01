package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"thedp.com/api/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	a := api.App{}
	a.Initialize(
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
	)

	a.Run(":80")
}
