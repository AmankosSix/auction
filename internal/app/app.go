package app

import (
	"auction/internal/config"
	"fmt"
	"log"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
