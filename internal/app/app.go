package app

import (
	"auction/internal/config"
	delivery "auction/internal/handler/http"
	"auction/internal/repository"
	"auction/internal/server"
	"auction/internal/service"
	"auction/pkg/auth"
	"auction/pkg/database"
	"auction/pkg/hash"
	"context"
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := database.NewClient(&cfg.Postgres)
	if err != nil {
		logrus.Fatal(err)
	}

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepositories(db)
	services := service.NewService(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
	})
	handlers := delivery.NewHandler(services)

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
