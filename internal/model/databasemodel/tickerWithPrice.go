package databasemodel

import "gorm.io/gorm"

type TickerWithPrice struct {
	gorm.Model
	Id     int     `gorm:"primaryKey; autoIncrement"`
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}
