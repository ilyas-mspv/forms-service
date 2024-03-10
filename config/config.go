package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	DB   `yaml:"db"`
}

type App struct {
	Name    string `yaml:"name" env:"APP_NAME" env-required:"true"`
	Version string `yaml:"version" env:"APP_VERSION" env-required:"true"`
}

type HTTP struct {
	Port int `yaml:"port" env-required:"true" env:"HTTP_PORT"`
}

type DB struct {
	ConnectionURL string `yaml:"url" env-required:"true" env:"DB_CONNECTION_URL"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
