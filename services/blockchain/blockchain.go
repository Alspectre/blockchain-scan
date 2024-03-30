package blockchain

import (
	"encoding/json"
	"fmt"
	"goblock/db/models"
	"goblock/services/handlers"
	"log"

	"gorm.io/gorm"
)

// NewBlockchainService creates a new BlockchainService instance
func NewBlockchainService(blockchain models.Blockchain, db *gorm.DB) *BlockchainService {
	service := &BlockchainService{
		blockchain: blockchain,
	}

	blockchainCurrencies, _ := models.GetActiveBlockchainCurrencies(db, blockchain.Key)
	service.blockchainCurrencies = blockchainCurrencies
	for _, currency := range blockchainCurrencies {
		service.currencies = append(service.currencies, currency.CurrencyId)
	}

	handler := handlers.Handler(blockchain.Client, blockchainCurrencies)
	if handler == nil {
		fmt.Printf("handling blockchain key %s error", blockchain.Key)
	}

	service.adapter = *handler

	return service
}

func (service *BlockchainService) LatestBlockNumber() (int64, error) {
	fetch, err := service.adapter.LatestBlockNumber(service.blockchain.Server)
	if err != nil {
		return 0, err
	}

	return fetch, nil
}

func (service *BlockchainService) Fetch(height int) {
	request, err := service.adapter.FetchBlock(service.blockchain.Server, height)
	if err != nil {
		log.Println(err)
	}

	jsonData, err := json.MarshalIndent(request, "", "    ")
	if err != nil {
		fmt.Println(jsonData)
		fmt.Println("Error:", err)
	}

}
