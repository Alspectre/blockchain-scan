package migrate

import (
	"fmt"
	"goblock/db"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	config       db.DatabaseConfig
	dbConnection *gorm.DB
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config = db.DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func Create() {
	Init()
	err := db.CreateDatabase(config, os.Getenv("DB_NAME"))
	fmt.Println(err)

	fmt.Println("--------- Create ---------")
}

func Migrate() {
	Init()
	dbConnection = db.ConfigDatabase(config, "")
	fmt.Println("--------- Migrate ---------")
}
