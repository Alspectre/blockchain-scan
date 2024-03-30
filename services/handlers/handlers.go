package handlers

import (
	"goblock/db/models"
	"goblock/lib/rpcclient"
)

func Handler(packageName string, bc []models.BlockchainCurrency) *rpcclient.ClientType {
	switch packageName {
	case "rpcclient":
		rpcClient := rpcclient.InitConfig(bc)

		return rpcClient
	}

	return &rpcclient.ClientType{}
}
