package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
		Auth     AuthConfig
	}

	PostgresConfig struct {
		DBName   string
		User     string
		Password string
		Host     string
		Port     string
		SSLMode  string
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}
)

func Init() (*Config, error) {
	var cfg Config

	initConfig()

	cfg.Postgres.setPostgresConfig()
	cfg.HTTP.setHTTPConfig()
	cfg.Auth.setAuthConfig()

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

func (c *HTTPConfig) setHTTPConfig() {
	c.Host = os.Getenv("HTTP_HOST")
	viper.UnmarshalKey("http", &c)
}

func (c *AuthConfig) setAuthConfig() {
	c.PasswordSalt = os.Getenv("PASSWORD_SALT")
	c.JWT.setJWTConfig()
}

func (c *JWTConfig) setJWTConfig() {
	c.SigningKey = os.Getenv("JWT_SIGNING_KEY")
	viper.UnmarshalKey("auth", &c)
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
