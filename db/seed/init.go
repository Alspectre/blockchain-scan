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
	dataMap    map[string][]map[string]interface{}
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
	// Blockchain()
	// Currencies()
	BlockchainCurrencies()
}

func Resource() {
	data, err := os.ReadFile("config/seed.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	if err := yaml.Unmarshal(data, &dataMap); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
}

func Blockchain() {
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

func Currencies() {
	currency := dataMap["currencies"]
	for _, b := range currency {
		currencies := models.Currencies{
			Name:      b["name"].(string),
			Precision: b["precision"].(int),
			IconUrl:   b["icon_url"].(string),
			MarketUrl: b["market_url"].(string),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := dbConnect.Create(&currencies).Error; err != nil {
			log.Fatalf("Error seeding blockchain: %v", err)
		}
		fmt.Printf("Seeded currencies: %+v\n", currencies)
	}
}

func BlockchainCurrencies() {
	blockService := dataMap["blockchain_currencies"]
	for _, b := range blockService {
		bc_service := models.BlockchainCurrency{
			CurrencyId:    b["currency_id"].(string),
			BlockchainKey: b["blockchain_key"].(string),
			BaseFactor:    b["base_factor"].(int),
			Status:        b["status"].(string),
			SmartContract: "",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := dbConnect.Create(&bc_service).Error; err != nil {
			log.Fatalf("Error seeding blockchain: %v", err)
		}
		fmt.Printf("Seeded currencies: %+v\n", currencies)
	}
}
