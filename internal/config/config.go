package config

import (
	"os"
)

type (
	Config struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		DBName   string
		User     string
		Password string
		Host     string
		Port     string
		SSLMode  string
	}
)

func Init() (*Config, error) {
	var cfg Config

	cfg.Postgres.setPostgresConfig()

	return &cfg, nil
}

func (p *PostgresConfig) setPostgresConfig() {
	p.DBName = os.Getenv("POSTGRES_DBNAME")
	p.User = os.Getenv("POSTGRES_USER")
	p.Password = os.Getenv("POSTGRES_PASS")
	p.Host = os.Getenv("POSTGRES_HOST")
	p.Port = os.Getenv("POSTGRES_PORT")
	p.SSLMode = os.Getenv("POSTGRES_SSL")
}
