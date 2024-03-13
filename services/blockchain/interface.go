package blockchain

import (
	"goblock/db/models"
	"goblock/lib/rpcclient"
	"sync"
)

type BlockchainService struct {
	blockchain             models.Blockchain
	blockchainCurrencies   []models.BlockchainCurrency
	currencies             []string
	adapter                rpcclient.ClientType
	latestBlockNumber      int64
	latestBlockNumberMutex sync.Mutex
}
