package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Logger
	Postgres
	Http
}

type Postgres struct {
	PostgresqlUrl string `mapstructure:"postgresql_url"`
}

type Logger struct {
	LogLevel string `mapstructure:"log_level"`
}

type Http struct {
	Port string `mapstucture:"port"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if os.Getenv("postgresql_url") != "" {
		config.Postgres.PostgresqlUrl = os.Getenv("postgresql_url")
	}

	return config, nil
}
