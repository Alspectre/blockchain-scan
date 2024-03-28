package main

import (
	"fmt"
	"goblock/db"
	"goblock/db/models"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

var (
	currencies models.Currencies
	blockchain models.Blockchain
	bc         models.BlockchainCurrency
	dbConnect  *gorm.DB
	err        error
)

func Init() {
	dbConnect, err = db.GetDatabaseConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func main() {
	Init()
	Resource()
}

func Resource() {
	data, err := os.ReadFile("config/seed.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var dataMap map[string][]map[string]interface{}
	if err := yaml.Unmarshal(data, &dataMap); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	blockchains := dataMap["blockchains"]
	for _, b := range blockchains {
		blockchain := models.Blockchain{
			Key:             b["key"].(string),
			Name:            b["name"].(string),
			Client:          b["client"].(string),
			Server:          b["server"].(string),
			Height:          b["height"].(int),
			Protocol:        b["protocol"].(string),
			MinConfirmation: b["min_confirmation"].(int),
			Status:          b["status"].(string),
			BlockchainGroup: b["blockchain_group"].(int),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if err := dbConnect.Create(&blockchain).Error; err != nil {
			log.Fatalf("Error seeding blockchain: %v", err)
		}
		fmt.Printf("Seeded blockchain: %+v\n", blockchain)
	}
}
