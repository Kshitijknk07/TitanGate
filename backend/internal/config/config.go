package config

import (
    "log"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found, using system environment variables")
    }
}