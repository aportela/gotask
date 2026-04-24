package configuration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init() {
	// TODO: return error (replace log.Fatal)
	viper.AddConfigPath(filepath.Join(".", "data"))
	viper.SetConfigName("gotask.config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Configuration file not found, creating new with default values")
			viper.SetDefault("database.path", "test")
			err := writeDefaultConfig()
			if err != nil {
				log.Fatal("Error creating configuration file:", err)
			}
		} else {
			log.Fatal("Error reading configuration file:", err)
		}
	}
	dbPath := viper.GetString("database.path")
	fmt.Println("Database path:", dbPath)
}

func writeDefaultConfig() error {
	databaseDir := filepath.Join(".", "data")
	if err := os.MkdirAll(databaseDir, os.ModePerm); err != nil {
		return err
	}

	viper.Set("database.path", "/data/gotask.sqlite3")

	configFile := filepath.Join(databaseDir, "gotask.config.yaml")
	err := viper.WriteConfigAs(configFile)
	if err != nil {
		return err
	}

	return nil
}
