package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/aviralbansal29/bill_split/constants"
)

// GlobalInstance should be reused.
type GlobalInstance struct {
	env        *viper.Viper
	database   *gorm.DB
	validation *validator.Validate
}

var globalInstance GlobalInstance

// InitiateGlobalInstance initiates the config data
func InitiateGlobalInstance() {
	env := loadConfig(constants.EnvVariablePath, constants.EnvFileName, constants.EnvFileExtension)
	database := loadDBHandler(env)
	globalInstance = GlobalInstance{
		env:        env,
		database:   database,
		validation: validator.New(),
	}
}

// GetEnv returns database instance
func GetEnv() *viper.Viper {
	return globalInstance.env
}

// DatabaseHandler returns postgres main database handler
func DatabaseHandler() *gorm.DB {
	return globalInstance.database
}

// Validator returns validator object
func Validator() *validator.Validate {
	return globalInstance.validation
}
