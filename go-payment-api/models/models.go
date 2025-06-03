package models

import (
	"gorm.io/gorm"
)


type user struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
}

type payment struct {
	gorm.model
	userID uint
	amount float64
	Currency    string
    Description string
    TransactionID string
}