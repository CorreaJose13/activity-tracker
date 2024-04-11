package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getValue(key, def string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return def
}
