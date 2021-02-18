package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Record struct {
	ID              int             `gorm:"primaryKey,column:id" json:"-"`
	FSymb           string          `gorm:"column:f_symb" json:"-"`
	TSymb           string          `gorm:"column:t_symb" json:"-"`
	Change24Hour    decimal.Decimal `gorm:"column:change_24_hour" json:"CHANGE24HOUR"`
	ChangePCT24Hour decimal.Decimal `gorm:"column:change_pct_24_hour" json:"CHANGEPCT24HOUR"`
	Open24Hour      decimal.Decimal `gorm:"column:open_24_hour" json:"OPEN24HOUR"`
	Volume24Hour    decimal.Decimal `gorm:"column:volume_24_hour" json:"VOLUME24HOUR"`
	Low24Hour       decimal.Decimal `gorm:"column:low_24_hour" json:"LOW24HOUR"`
	High24Hour      decimal.Decimal `gorm:"column:high_24_hour" json:"HIGH24HOUR"`
	Price           decimal.Decimal `gorm:"column:price" json:"PRICE"`
	Supply          decimal.Decimal `gorm:"column:supply" json:"SUPPLY"`
	MKTCAP          decimal.Decimal `gorm:"column:mktcap" json:"MKTCAP"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"-"`
}

func (r *Record) Table() string {
	return "records"
}
