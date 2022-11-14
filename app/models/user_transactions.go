package models

import (
	"github.com/aviralbansal29/bill_split/config"
	"gorm.io/gorm"
)

// UserTransaction defines relation between users owed amount for each transaction
type UserTransaction struct {
	gorm.Model

	TransactionID uint        `json:"transaction_id" gorm:"index"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`

	FromID uint `json:"from_id" gorm:"index"`
	From   User `gorm:"foreignKey:FromID;references:ID"`

	ToID uint `json:"to_id" gorm:"index"`
	To   User `gorm:"foreignKey:ToID;references:ID"`

	Amount  uint
	Pending int `gorm:"-"`
}

// CreateEntry creates entry in DB
func (t *UserTransaction) CreateEntry() error {
	return config.DatabaseHandler().Create(&t).Error
}

// BulkCreateUserTransaction creates entries in bulk
func BulkCreateUserTransaction(entries []UserTransaction) error {
	return config.DatabaseHandler().Create(&entries).Error
}
