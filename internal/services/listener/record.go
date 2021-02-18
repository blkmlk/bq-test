package listener

import "github.com/shopspring/decimal"

type Records map[string]map[string]*Record

type Record struct {
	Change24Hour    decimal.Decimal `json:"CHANGE24HOUR"`
	ChangePCT24Hour decimal.Decimal `json:"CHANGEPCT24HOUR"`
	Open24Hour      decimal.Decimal `json:"OPEN24HOUR"`
	Volume24Hour    decimal.Decimal `json:"VOLUME24HOUR"`
	Low24Hour       decimal.Decimal `json:"LOW24HOUR"`
	High24Hour      decimal.Decimal `json:"HIGH24HOUR"`
	Price           decimal.Decimal `json:"PRICE"`
	Supply          decimal.Decimal `json:"SUPPLY"`
	MKTCAP          decimal.Decimal `json:"MKTCAP"`
}
