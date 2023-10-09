package databasemodel

import "gorm.io/gorm"

type Ticker struct {
	gorm.Model
	Id     int    `gorm:"primaryKey; autoIncrement"`
	Ticker string `json:"ticker"`
}
