package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Config holds all the configuration variables
type Config struct {
	Server     ServerConfig `mapstructure:"server"`
	UserApiURL string       `mapstructure:"user_api_url"`
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// LoadConfig loads the configuration file using Viper
func LoadConfig() (*Config, error) {
	viper.SetDefault("server.port", 8080)

	// Set the configuration file name and path
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Check environment variable `CONFIG_PATH` to get the config path if provided
	configPath := "tests" // Default path
	if envConfigPath := viper.GetString("CONFIG_PATH"); envConfigPath != "" {
		configPath = envConfigPath
	}

	viper.AddConfigPath(configPath) // Path to look for the config file

	// Enable reading configuration from environment variables
	viper.AutomaticEnv()

	// Read in the configuration file if available
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Unable to read config file, using defaults and environment variables: %v", err)
	}

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}
