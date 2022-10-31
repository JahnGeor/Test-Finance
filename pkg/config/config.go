package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "./configs/.env"
)

type Config struct {
	DBConfig
	ServerConfig
}

type DBConfig struct {
	Username    string `env:"DB_USERNAME"`
	Password    string `env:"DB_PASSWORD"`
	Host        string `env:"DB_HOST"`
	Port        string `env:"DB_PORT"`
	DBName      string `env:"DB_NAME"`
	SSLmode     string `env:"DB_SSLMODE"`
	MaxAttempts int    `env:"DB_MAXATTEMPTS"`
}

type ServerConfig struct {
	Port string `env:"SRV_PORT"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
