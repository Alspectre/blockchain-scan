package models

import (
	"github.com/hashicorp/vault/api"
	"gorm.io/gorm"
)

type BlockchainCurrency struct {
	CurrencyId            string     `json:"currency_id" db:"currency_id"`
	BlockchainKey         string     `json:"blockchain_key" db:"blockchain_key"`
	ParentID              string     `json:"parent_id" db:"parent_id"`
	DepositFee            float64    `json:"deposit_fee" db:"deposit_fee"`
	MinDepositAmount      float64    `json:"min_deposit_amount" db:"min_deposit_amount"`
	MinCollectionAmount   float64    `json:"min_collection_amount" db:"min_collection_amount"`
	WithdrawFee           float64    `json:"withdraw_fee" db:"withdraw_fee"`
	MinWithdrawAmount     float64    `json:"min_withdraw_amount" db:"min_withdraw_amount"`
	DepositEnabled        bool       `json:"deposit_enabled" db:"deposit_enabled"`
	WithdrawalEnabled     bool       `json:"withdrawal_enabled" db:"withdrawal_enabled"`
	AutoUpdateFeesEnabled bool       `json:"auto_update_fees_enabled" db:"auto_update_fees_enabled"`
	BaseFactor            int64      `json:"base_factor" db:"base_factor"`
	Status                string     `json:"status" db:"status"`
	Options               []byte     `gorm:"type:json" db:"options"`
	Blockchain            Blockchain `gorm:"foreignKey:BlockchainKey;references:key"`
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

func GetActiveBlockchainCurrencies(db *gorm.DB, vault *api.Client) ([]BlockchainCurrency, error) {
	var blockchain_currencies []BlockchainCurrency
	if err := db.Preload("Blockchain").Joins("JOIN blockchains ON blockchains.key = blockchain_currencies.blockchain_key").
		Where("blockchains.status = ?", "active").
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
