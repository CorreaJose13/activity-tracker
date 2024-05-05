package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Loading .env file failed with '%s'\n", err)
	}
}

func getValue(key, def string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return def
}
