package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		User     string
		Password string
		Host     string
		Port     string
	}
)

func Init() (*Config, error) {
	var cfg Config

	if err := godotenv.Load("../.env"); err != nil {
		return nil, err
	}

	cfg.Postgres.setPostgresConfig()

	return &cfg, nil
}

func (p *PostgresConfig) setPostgresConfig() {
	p.User = os.Getenv("POSTGRES_USER")
	p.Password = os.Getenv("POSTGRES_PASS")
	p.Host = os.Getenv("POSTGRES_HOST")
	p.Port = os.Getenv("POSTGRES_PORT")
}
