package factory

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile("../migration.yaml")
	readConfig()
}

func readConfig() {
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("migrator config file not found: %w", err))
		}
		panic(fmt.Errorf("viper fatal error: %w", err))
	}
}
