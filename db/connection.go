package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
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
				dataSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.Username, dbConfig.Password, Database, dbConfig.Port)
				checkDatabaseExists(Database, config)
			} else {
				dataSource = fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Port)
				checkDatabaseExists(Database, config)
			}

			dbConnection, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
			if err != nil {
				log.Fatalf("%v", err)
				return
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
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	connect := ConfigDatabase(config, os.Getenv("DB_NAME"))

	return connect, err
}

func checkDatabaseExists(dbName string, config DatabaseConfig) (bool, error) {
	dataSource = fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.Username, config.Password, config.Port)
	db, err := sql.Open(config.Driver, dataSource)
	if err != nil {
		err = fmt.Errorf("error connecting to database: %w", err)
		return false, err
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", dbName)
	var exists bool

	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("database '%s' does not exist", dbName)
		return false, err
	}

	return exists, nil
}

func CreateDatabase(config DatabaseConfig, dbName string) error {
	dataSource = fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.Username, config.Password, config.Port)
	db, err := sql.Open(config.Driver, dataSource)
	if err != nil {
		err = fmt.Errorf("error connecting to database: %w", err)
		return err
	}
	defer db.Close()

	createStmt := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err = db.Exec(createStmt)
	if err != nil {
		return err
	}

	return err
}
