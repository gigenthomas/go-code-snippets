package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}

// ConfigProvider is an interface that defines methods to get configuration values.
// Enables validation to establish a contract for configuration providers.
// This is useful for testing and mocking configuration providers.
type ConfigProvider interface {
	GetPort() int
	GetName() string
}

func LoadConfig() (*Config, error) {
	// Initialize viper to read from conf.env file

	rootDir, err := os.Getwd()
	fmt.Println("Root Directory:", rootDir)
	if err != nil {
		return nil, fmt.Errorf("unable to get working directory: %w", err)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("conf")

	envPath := filepath.Join(rootDir, "config.yaml")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %s: %w", envPath, err)
	}

	viper.SetConfigFile(envPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(viper.GetString("PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: port,
		Name: viper.GetString("NAME"),
	}, nil
}

func (c *Config) GetPort() int {
	return c.Port
}

func (c *Config) GetName() string {
	return c.Name
}
