package handlers

import (
	"goblock/db/models"
	"goblock/lib/rpcclient"
)

func Handler(packageName string, bc []models.BlockchainCurrency) *rpcclient.ClientType {
	switch packageName {
	case "rpc":
		rpcClient := rpcclient.InitConfig(bc)

		return rpcClient
	}

	return &rpcclient.ClientType{}
}
