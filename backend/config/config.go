package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Name string
	Host string
	Port int
}

type Specialurl struct {
	UrlWithoutToken []string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./config/config.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
		return nil, err
	}

	return &config, nil
}
