package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/aviralbansal29/bill_split/config"
	"gorm.io/gorm"
)

// Transaction defines structure for storing transactions in a group
type Transaction struct {
	gorm.Model

	GroupID uint  `json:"group_id" gorm:"index"`
	Group   Group `gorm:"foreignKey:GroupID;references:ID"`

	Amount uint
	Name   string
}

func (t Transaction) validateAmount() error {
	if t.Amount <= 0 {
		return errors.New("Amount must be greater than 0")
	}
	return nil
}

func (t Transaction) validate() error {
	return t.validateAmount()
}

// CreateEntry creates entry in DB
func (t *Transaction) CreateEntry() error {
	err := t.validate()
	if err != nil {
		return err
	}

	return config.DatabaseHandler().Create(&t).Error
}

// DeleteEntry soft deletes entry in DB
func (t *Transaction) DeleteEntry() error {
	return config.DatabaseHandler().Delete(&t).Error
}

// GetListQuery returns list query based on user-id
func (t Transaction) GetListQuery(userID uint) *gorm.DB {
	query := config.DatabaseHandler().Model(&Transaction{}).Joins(
		"inner join user_transactions on transactions.id = user_transactions.transaction_id",
	).Joins(
		"inner join groups on groups.id = transactions.group_id",
	).Where(
		"user_transactions.from_id = ? OR user_transactions.to_id = ?", userID, userID,
	).Select(
		"transactions.created_at as date, split_part(string_agg(groups.name, ','), ',', 1) as group, transactions.name as expense",
		"transactions.amount as amount",
		fmt.Sprintf(
			"sum(case when user_transactions.from_id=%d then user_transactions.amount*-1 else user_transactions.amount end) as pending",
			userID,
		),
	).Group("transactions.id")
	return query
}

func (t Transaction) AddGroupFilter(query *gorm.DB, groupID uint) *gorm.DB {
	return query.Where("groups.id = ?", groupID)
}

func (t Transaction) AddTimePeriodFilter(query *gorm.DB, from time.Time, to time.Time) *gorm.DB {
	return query.Where("transactions.created_at > ? and transactions.created_at < ?", from, to)
}
