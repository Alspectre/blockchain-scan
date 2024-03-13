package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

var (
	dbConfig     DatabaseConfig
	dbConnection *gorm.DB
	once         sync.Once
	err          error
	dataSource   string
)

func ConfigDatabase(config DatabaseConfig, Database string) *gorm.DB {
	dbConfig = config
	once.Do(
		func() {
			if Database != "" {
				dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, Database)
			} else {
				dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
			}
			dbConnection, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
			if err != nil {
				log.Fatalf("Error connecting to database: %v", err)
			}

			fmt.Println("Connected to the database!")
		})

	return dbConnection
}

func GetDatabaseConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config := DatabaseConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	connect := ConfigDatabase(config, os.Getenv("DB_NAME"))

	return connect, err
}
