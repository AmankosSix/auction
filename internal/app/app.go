package app

import (
	"auction/internal/config"
	"auction/pkg/database"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.NewClient(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
}
