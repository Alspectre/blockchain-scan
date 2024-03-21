package migrate

import (
	"bytes"
	"fmt"
	"goblock/db"
	"log"
	"os"
	"os/exec"

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
	fmt.Println("--------- Migrate ---------")
	databaseSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, os.Getenv("DB_NAME"))
	cmd := exec.Command("migrate", "-path", "db/migration", "-database", databaseSource, "-verbose", "up")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(out)
	}
}
