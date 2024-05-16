package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProxyHost   string
	BaseURL     string
	TwoURL      string
	UserAgent   string
	City        string
	Street      string
	HouseNumber string
}

func LoadConfig() (*Config, error) {
	// Загрузка файла .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Инициализация конфига
	cfg := &Config{
		ProxyHost:   os.Getenv("PROXY_HOST"),
		BaseURL:     os.Getenv("BASE_URL"),
		TwoURL:      os.Getenv("TWO_URL"),
		UserAgent:   os.Getenv("USER_AGENT"),
		City:        os.Getenv("CITY"),
		Street:      os.Getenv("STREET"),
		HouseNumber: os.Getenv("NUMBER_HOUSE"),
	}

	return cfg, nil
}
