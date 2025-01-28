package mirrorMove

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Settings struct {
    DB_CONNECTION  string
}

func GetENV() Settings {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := os.Getenv("DB_CONNECTION")

	settings := Settings {
		dbConnection,
	}

	return settings
}