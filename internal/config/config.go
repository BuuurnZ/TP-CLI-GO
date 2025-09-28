package config

import (
	"fmt"
	"loganalyzer/internal/stores"

	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		Type     string `mapstructure:"type"`
		Database string `mapstructure:"database"`
		JSONFile string `mapstructure:"json_file"`
	} `mapstructure:"storage"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("storage.type", "memory")
	viper.SetDefault("storage.database", "contacts.db")
	viper.SetDefault("storage.json_file", "contacts.json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewStore(config *Config) (stores.Storer, error) {
	switch config.Storage.Type {
	case "gorm":
		return stores.NewGORMStore(config.Storage.Database)
	case "json":
		return stores.NewJSONStore(config.Storage.JSONFile)
	case "memory":
		return stores.NewMemoryStore(), nil
	default:
		return nil, fmt.Errorf("type de stockage non supporté: %s", config.Storage.Type)
	}
}