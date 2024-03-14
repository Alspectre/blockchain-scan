package blockchain

import (
	"goblock/db/models"
	"goblock/lib/rpcclient"
)

type BlockchainService struct {
	blockchain           models.Blockchain
	blockchainCurrencies []models.BlockchainCurrency
	currencies           []string
	adapter              rpcclient.ClientType
}
