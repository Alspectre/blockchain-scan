package migrate

import (
	"fmt"
	"goblock/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config := db.DatabaseConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	dbConfig := db.ConfigDatabase(config, "")
	fmt.Println(dbConfig)
}
