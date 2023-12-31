package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	a := App{}
	a.Initialize(
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
	)

	a.Run(":80")
}
