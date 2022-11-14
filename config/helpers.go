package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// projectRootPath gets the current_project root-path
func projectRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	// remove the current-folder from the current-directory of file
	rootPath := strings.Split(filepath.Dir(b), "/config")
	return rootPath[0]
}

// loadConfig loads configurations for project
func loadConfig(path string, filename string, fileExtension string) *viper.Viper {
	// load env file
	environment := viper.New()
	environment.AddConfigPath(projectRootPath() + path)
	environment.SetConfigName(filename)
	environment.SetConfigType(fileExtension)
	environment.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	environment.AutomaticEnv()
	environment.WatchConfig()
	err := environment.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return environment
}

func loadDBHandler(env *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.GetString("postgres.host"),
		env.GetString("postgres.user"),
		env.GetString("postgres.password"),
		env.GetString("postgres.db"),
		env.GetString("postgres.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
