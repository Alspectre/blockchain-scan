package main

import (
	"fmt"
	"goblock/config/vaultconfig"
	"goblock/db"
	"goblock/db/models"
	"goblock/services/blockchain"
	"goblock/services/handlers"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	run    bool = true
	params string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	testVault, err := vaultconfig.InitVault()
	if err != nil {
		fmt.Println(err)
	}

	vaultToken := os.Getenv("VAULT_TOKEN")

	testVault.SetToken(vaultToken)

	connectionDb, err := db.GetDatabaseConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Printf("Connected to Database")

	log.Printf("Try to get all data blockchain")

	for run {
		blockchains, err := models.GetAllBlockchainWitHDecrypt(connectionDb, testVault, "active")
		if err != nil {
			log.Fatalf("Error retrieving blockchain data: %v", err)
		}
		for _, v := range blockchains {
			fmt.Printf("Start fetching %s on block %d", v.Key, v.Height)
			fmt.Println("")

			if v.Server == "" {
				fmt.Println("Server is null on blockchain currencies %s", v.Key)
				fmt.Println(err)
				continue
			}

			params = fmt.Sprintf("blockchain_key = '%s'", v.Key)
			bcCurrencies, _ := models.GetFilteredBlockchainChurrencies(connectionDb, params)

			handler := handlers.Handler(v.Client, bcCurrencies)
			if handler == nil {
				continue
			}

			bc_service := blockchain.NewBlockchainService(v, connectionDb, testVault, handler)
			fmt.Printf("Fetching block %s on height %d with server %s", v.Client, v.Height, v.Server)
			fmt.Println("")
			lastBlock, err := bc_service.LatestBlockNumber()
			if err != nil {
				fmt.Println("Error:", err)
			}

			if int64(v.Height) < lastBlock-3 {
				bc_service.Fetch(v.Height)
			}
			fmt.Printf("Done fetching %s on block %d", v.Key, v.Height)
			fmt.Println("")
			models.UpdateHeight(connectionDb, testVault, int(v.ID), int(v.Height+1))
			fmt.Println("")
		}
		time.Sleep(2 * time.Second)
	}

	if !run {
		fmt.Printf("Blockchain Service has Stopped")
	}
}
