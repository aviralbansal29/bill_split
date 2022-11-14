package models

import (
	"errors"
	"strconv"

	"github.com/aviralbansal29/bill_split/config"
	"gorm.io/gorm"
)

// User defines structure of user data stored
type User struct {
	gorm.Model

	Name        string
	PhoneNumber string `json:"phone_number" gorm:"unique;index"`
}

func (u User) validatePhoneNumber() error {
	_, err := strconv.ParseUint(u.PhoneNumber, 10, 64)
	if err != nil {
		return errors.New("Phone Number should be numeric")
	}
	if len(u.PhoneNumber) != 10 {
		return errors.New("Phone number length should be 10")
	}
	var count int64
	config.DatabaseHandler().Model(&u).Where("phone_number = ?", u.PhoneNumber).Count(&count)
	if count != 0 {
		return errors.New("Account with same phone number already exists")
	}
	return nil
}

func (u User) validateName() error {
	if len(u.Name) == 0 {
		return errors.New("Name length should be more than 0")
	}
	return nil
}

func (u User) validate() error {
	err := u.validateName()
	if err != nil {
		return err
	}

	err = u.validatePhoneNumber()
	if err != nil {
		return err
	}

	return nil
}

// CreateEntry creates user entry in DB
func (u *User) CreateEntry() error {
	err := u.validate()
	if err != nil {
		return err
	}

	err = config.DatabaseHandler().Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
