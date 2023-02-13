package app

import (
	"auction/internal/config"
	delivery "auction/internal/handler/http"
	"auction/internal/server"
	"auction/pkg/database"
	"context"
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	handlers := delivery.NewHandler()

	srv := server.NewServer(&cfg.HTTP, handlers.Init(cfg))

	go func() {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Stop(ctx); err != nil {
		logrus.Errorf("failed to stop server: %v", err)
	}
}
