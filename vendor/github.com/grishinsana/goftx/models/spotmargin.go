package models

import "github.com/shopspring/decimal"

type LendingInfo struct {
	Coin     string          `json:"coin"`
	Lendable decimal.Decimal `json:"lendable"`
	Locked   decimal.Decimal `json:"locked"`
	MinRate  decimal.Decimal `json:"minRate"`
	Offered  decimal.Decimal `json:"offered"`
}

type LendingRate struct {
	Coin     string          `json:"coin"`
	Estimate decimal.Decimal `json:"estimate"`
	Previous decimal.Decimal `json:"previous"`
}
