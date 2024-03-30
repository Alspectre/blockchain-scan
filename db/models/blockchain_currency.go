package models

import (
	"time"

	"gorm.io/gorm"
)

type BlockchainCurrency struct {
	CurrencyId    string     `json:"currency_id" db:"currency_id"`
	BlockchainKey string     `json:"blockchain_key" db:"blockchain_key"`
	ParentID      string     `json:"parent_id" db:"parent_id"`
	BaseFactor    int        `json:"base_factor" db:"base_factor"`
	Status        string     `json:"status" db:"status"`
	Options       []byte     `gorm:"type:json" db:"options"`
	SmartContract string     `json:"smart_contract" db:"smart_contract"`
	Blockchain    Blockchain `gorm:"foreignKey:BlockchainKey;references:key"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

type Option struct {
	GasPrice             string `omitempty,json:"gas_price"`
	GasLimit             string `omitempty,json:"gas_limit"`
	Erc20ContractAddress string `omitempty,json:"erc20_contract_address"`
}

type BlockchainKey struct {
	BlockchainKey string `omitempty,json:"blockchain_key" gorm:"primaryKey"`
}

func GetAllBlockchainCurrencies(db *gorm.DB) ([]BlockchainCurrency, error) {
	var blockchain_currencies []BlockchainCurrency
	if err := db.Preload("Blockchain").Find(&blockchain_currencies).Error; err != nil {
		return nil, err
	}
	return blockchain_currencies, nil
}

func GetActiveBlockchainCurrencies(db *gorm.DB, key string) ([]BlockchainCurrency, error) {
	var blockchain_currencies []BlockchainCurrency
	if err := db.Preload("Blockchain").Joins("JOIN blockchains ON blockchains.key = blockchain_currencies.blockchain_key").
		Where("blockchains.status = ?", "active").Where("blockchain_currencies.blockchain_key = ?", key).
		Find(&blockchain_currencies).Error; err != nil {
		return nil, err
	}

	return blockchain_currencies, nil
}

func GetFilteredBlockchainChurrencies(db *gorm.DB, params string) ([]BlockchainCurrency, error) {
	var blockchain_currencies []BlockchainCurrency
	if err := db.Where(params).Find(&blockchain_currencies).Error; err != nil {
		return nil, err
	}

	return blockchain_currencies, nil
}
