package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Task2API/config"
	"Task2API/internal/handler"
	"Task2API/internal/middleware"
	"Task2API/internal/repository"
	"Task2API/internal/routes"
	"Task2API/internal/services"
	"Task2API/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.MustLoadConfig("config.yml")

	log := logger.NewStdLogger()
	logger.ConfigureFileLogger(log,
		cfg.Logging.Filename,
		cfg.Logging.MaxSizeMB,
		cfg.Logging.MaxBackups,
		cfg.Logging.MaxAgeDays,
	)

	dbPool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Error("Unable to create connection pool: %v", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	tokenRepo := repository.NewTokenRepository(dbPool, log)

	rateLimitMiddleware := middleware.NewRateLimitMiddleware(cfg.RateLimit.RequestsPerSecond, log)

	targetAPIConfig := services.TargetAPIConfig{
		URL:           cfg.TargetAPI.URL,
		Authorization: cfg.TargetAPI.Authorization,
	}
	predictionService := services.NewPredictionService(targetAPIConfig, log)

	h := handler.NewHandler(predictionService, log)

	authMiddleware := middleware.NewAuthMiddleware(tokenRepo, log)

	router := routes.NewRouter(h, authMiddleware, rateLimitMiddleware)
	handlerChain := router.SetupRoutes()

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      handlerChain,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting API gateway on :%s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Failed to start server: %v", err)
		}
	}()

	<-done
	log.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Failed to shutdown server gracefully: %v", err)
	}
	log.Info("Server stopped")
}
