package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	File string
)

func Initiate() {
	if File != "" {
		viper.SetConfigFile(File)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; Use default values
		} else {
			panic(fmt.Errorf("Can't parse config: %s\n", err))
		}
	}
}
