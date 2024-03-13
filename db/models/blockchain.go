package models

import (
	"fmt"
	"goblock/config/vaultconfig"

	"github.com/hashicorp/vault/api"
	"gorm.io/gorm"
)

type Blockchain struct {
	ID                 int64   `json:"id" db:"id" gorm:"primaryKey"`
	Key                string  `json:"key" db:"key" `
	Name               string  `json:"name" db:"name"`
	Client             string  `json:"client" db:"client"`
	Server             string  `omitempty,json:"server"`
	ServerEncrypted    string  `json:"server_encrypted" db:"server_encrypted"`
	Height             int     `json:"height" db:"height"`
	CollectionGasSpeed string  `json:"collection_gas_speed" db:"collection_gas_speed"`
	WithdrawalGasSpeed string  `json:"withdrawal_gas_speed" db:"withdrawal_gas_speed"`
	Protocol           string  `json:"protocol" db:"protocol"`
	MinConfirmation    int8    `json:"min_confirmation" db:"min_confirmation"`
	MinDepositAmount   float64 `json:"min_deposit_amount" db:"min_deposit_amount"`
	WithdrawFee        float64 `json:"withdraw_fee" db:"withdraw_fee"`
	MinWithdrawAmount  float64 `json:"min_withdraw_amount" db:"min_withdraw_amount"`
	Status             string  `json:"status" db:"status"`
	BlockchainGroup    int     `json:"blockchain_group" db:"blockchain_group"`
}

type Blockchains []Blockchain

func GetAllBlockchain(db *gorm.DB) ([]Blockchain, error) {
	var blockchains Blockchains
	_ = db.Find(&blockchains)

	return blockchains, nil
}

func GetAllBlockchainWitHDecrypt(db *gorm.DB, client *api.Client, status string) ([]Blockchain, error) {
	var blockchains Blockchains
	_ = db.Where("status = ?", status).Find(&blockchains)

	for i, v := range blockchains {
		decrypt, _ := vaultconfig.DecryptValue(client, v.ServerEncrypted)
		blockchains[i].Server = decrypt
	}

	return blockchains, nil
}

func UpdateHeight(db *gorm.DB, client *api.Client, id int, height int) string {
	var blockchain Blockchain
	if err := db.First(&blockchain, id).Error; err != nil {
		fmt.Printf("Failed to find blockchain with ID %d: %v\n", id, err)
	}

	if err := db.Model(&blockchain).Update("height", height); err != nil {
		server, _ := vaultconfig.DecryptValue(client, blockchain.ServerEncrypted)
		fmt.Printf("Blockchain %s updated on hight %d", server, blockchain.Height)
	}
	return string(blockchain.Height)
}
