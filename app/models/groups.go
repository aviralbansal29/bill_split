package models

import (
	"errors"

	"github.com/aviralbansal29/bill_split/config"
	"gorm.io/gorm"
)

// Group defines structure of Group
type Group struct {
	gorm.Model

	Name      string `gorm:"index,unique"`
	IsSettled bool   `json:"is_settled"`

	UserGroups []UserGroup `gorm:"foreignKey:GroupID" json:"user_groups"`
}

func (g Group) validateName() error {
	if len(g.Name) == 0 {
		return errors.New("Name length should be more than 0")
	}

	var count int64
	config.DatabaseHandler().Model(&g).Where("name = ?", g.Name).Count(&count)
	if count > 0 {
		return errors.New("Group with same name already exists")
	}
	return nil
}

func (g Group) validate() error {
	err := g.validateName()
	if err != nil {
		return err
	}
	return nil
}

// CreateEntry creates group entry in DB
func (g *Group) CreateEntry() error {
	err := g.validate()
	if err != nil {
		return err
	}
	g.IsSettled = false

	return config.DatabaseHandler().Create(&g).Error
}

// DeleteEntry deletes entry
func (g *Group) DeleteEntry() error {
	return config.DatabaseHandler().Delete(&g).Error
}
