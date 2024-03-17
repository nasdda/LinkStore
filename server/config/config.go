package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress         string `mapstructure:"server_address"`
	MongoConnectionString string `mapstructure:"mongo_connection_string"`
	MongoDBName           string `mapstructure:"mongo_db_name"`
	Port                  string `mapstructure:"port"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.SetDefault("server_address", "localhost:8080")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &cfg, nil
}
