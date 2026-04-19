package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Enable automatic environment variable binding for server deployment
	v.AutomaticEnv()

	// Read .env file for local development
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		// Ignore error if config file is not found
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct")
	}

	return &config, nil
}
