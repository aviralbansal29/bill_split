package models

import (
	"errors"

	"github.com/aviralbansal29/bill_split/config"
	"gorm.io/gorm"
)

// UserGroup defines relationship between user and group
type UserGroup struct {
	gorm.Model

	UserID uint `json:"user_id" gorm:"index:idx_user_group,unique"`
	User   User `gorm:"foreignKey:UserID;references:ID"`

	GroupID uint  `json:"group_id" gorm:"index:idx_user_group,unique"`
	Group   Group `gorm:"foreignKey:GroupID;references:ID"`

	Owed uint
}

func (g UserGroup) validate() error {
	var count int64
	config.DatabaseHandler().Model(&g).Where("user_id = ? and group_id = ?", g.UserID, g.GroupID).Count(&count)
	if count > 0 {
		return errors.New("User Already added to group")
	}
	return nil
}

// CreateEntry creates user-group entry in DB
func (g *UserGroup) CreateEntry() error {
	err := g.validate()
	if err != nil {
		return err
	}

	return config.DatabaseHandler().Create(&g).Error
}

func BulkCreateGroups(grps []UserGroup) error {
	return config.DatabaseHandler().Create(&grps).Error
}
