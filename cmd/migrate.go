package main

import (
	"auction/internal/config"
	"auction/internal/schema"
	"auction/pkg/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err = schema.Migrator(db.DB); err != nil {
		log.Fatal(err)
	}
}
