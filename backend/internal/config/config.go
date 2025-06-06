package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	paths := []string{".env", "../.env", "../../.env"}
	loaded := false
	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			loaded = true
			break
		}
	}
	if !loaded {
		log.Printf("Warning: .env file not found, using system environment variables")
	}
}
