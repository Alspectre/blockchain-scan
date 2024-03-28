package models

type Currencies struct {
	ID        int64  `json:"id" db:"id" gorm:"primaryKey"`
	Name      string `json:"name" db:"name" `
	Precision string `omitempty,json:"precision" db:"precision"`
	IconUrl   string `omitempty,json:"icon_url" db:"icon_url"`
	MarketUrl string `omitempty,json:"market_url" db:"market_url"`
}
