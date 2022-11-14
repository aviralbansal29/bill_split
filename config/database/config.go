package database

import (
	"log"

	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/aviralbansal29/bill_split/config"
)

// Migrate migrates the DB
func Migrate() {
	if err := config.DatabaseHandler().AutoMigrate(
		&models.User{}, &models.Group{}, &models.UserGroup{}, &models.Transaction{}, &models.UserTransaction{},
	); err != nil {
		log.Fatal(err)
	}
}
