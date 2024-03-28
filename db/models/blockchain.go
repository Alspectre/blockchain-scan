package models

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
	"gorm.io/gorm"
)

type Blockchain struct {
	ID              int64     `json:"id" db:"id" gorm:"primaryKey"`
	Key             string    `json:"key" db:"key" `
	Name            string    `json:"name" db:"name"`
	Client          string    `json:"client" db:"client"`
	Server          string    `omitempty,json:"server"`
	Height          int       `json:"height" db:"height"`
	Protocol        string    `json:"protocol" db:"protocol"`
	MinConfirmation int       `json:"min_confirmation" db:"min_confirmation"`
	Status          string    `json:"status" db:"status"`
	BlockchainGroup int       `json:"blockchain_group" db:"blockchain_group"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Blockchains []Blockchain

func CreateBlockchain(db *gorm.DB, blockchain *Blockchain) error {
	if err := db.Create(blockchain).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBlockchain(db *gorm.DB) ([]Blockchain, error) {
	var blockchains Blockchains
	_ = db.Find(&blockchains)

	return blockchains, nil
}

func GetAllBlockchainWitHDecrypt(db *gorm.DB, client *api.Client, status string) ([]Blockchain, error) {
	var blockchains Blockchains
	_ = db.Where("status = ?", status).Find(&blockchains)

	return blockchains, nil
}

func UpdateHeight(db *gorm.DB, client *api.Client, id int, height int) string {
	var blockchain Blockchain
	if err := db.First(&blockchain, id).Error; err != nil {
		fmt.Printf("Failed to find blockchain with ID %d: %v\n", id, err)
	}

	if err := db.Model(&blockchain).Update("height", height); err != nil {
		fmt.Printf("Blockchain updated on hight %d", blockchain.Height)
	}
	return string(blockchain.Height)
}
