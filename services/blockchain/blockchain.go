package blockchain

import (
	"encoding/json"
	"fmt"
	"goblock/db/models"
	"goblock/lib/rpcclient"
	"log"

	"github.com/hashicorp/vault/api"
	"gorm.io/gorm"
)

// NewBlockchainService creates a new BlockchainService instance
func NewBlockchainService(blockchain models.Blockchain, db *gorm.DB, vault *api.Client, handler *rpcclient.ClientType) *BlockchainService {
	service := &BlockchainService{
		blockchain: blockchain,
	}

	blockchainCurrencies, _ := models.GetActiveBlockchainCurrencies(db, vault)
	service.blockchainCurrencies = blockchainCurrencies
	for _, currency := range blockchainCurrencies {
		service.currencies = append(service.currencies, currency.CurrencyId)
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
		fmt.Println("Error:", err)
	}

	fmt.Println(string(jsonData))

}
