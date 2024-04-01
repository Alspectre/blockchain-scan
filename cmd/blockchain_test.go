package main

import (
	"fmt"
	"goblock/db"
	"goblock/db/models"
	"goblock/services/blockchain"
	"log"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

var (
	running bool = true
)

func mainTest() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	connectionDb, err := db.GetDatabaseConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Printf("Connected to Database")

	log.Printf("Try to get all data blockchain")

	for running {
		blockchains, err := models.GetAllBlockchainWitHDecrypt(connectionDb, "enabled")
		if err != nil {
			log.Fatalf("Error retrieving blockchain data: %v", err)
		}
		for _, v := range blockchains {
			fmt.Printf("Start fetching %s on block %d", v.Key, v.Height)
			fmt.Println("")

			if v.Server == "" {
				fmt.Printf("Server is null on blockchain currencies %s", v.Key)
				fmt.Println(err)
				continue
			}

			bc_service := blockchain.NewBlockchainService(v, connectionDb)
			fmt.Printf("Fetching block %s on height %d", v.Client, v.Height)
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
			models.UpdateHeight(connectionDb, int(v.ID), int(v.Height+1))
			fmt.Println("")
		}
		// time.Sleep(1 * time.Second)
	}

	if !running {
		fmt.Printf("Blockchain Service has Stopped")
	}
}

func BenchmarkMemoryConsumption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Perform GC to get accurate memory consumption
		runtime.GC()
		var m0, m1 runtime.MemStats
		runtime.ReadMemStats(&m0)

		// Run the main function
		mainTest()

		// Read memory stats again after running main function
		runtime.ReadMemStats(&m1)

		// Calculate memory consumed
		memConsumed := m1.TotalAlloc - m0.TotalAlloc
		b.ReportMetric(float64(memConsumed), "bytes")
	}
}
