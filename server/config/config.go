package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	DatabaseURL   string `mapstructure:"database_url"`
	MongoURI      string `mapstructure:"mongo_uri"`
}

func LoadConfig(configPaths ...string) (*Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.SetDefault("server_address", "localhost:8080")

	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &cfg, nil
}

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println("Config loaded:", config)
}
