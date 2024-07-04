package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BaseURL    string
	CustomerID string
	Key        string
)

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BaseURL = os.Getenv("BASE_URL")
	CustomerID = os.Getenv("CUSTOMER_ID")
	Key = os.Getenv("KEY")

	if BaseURL == "" || CustomerID == "" || Key == "" {
		log.Fatal("Environment variables BASE_URL, CUSTOMER_ID, and KEY must be set")
	}
}
