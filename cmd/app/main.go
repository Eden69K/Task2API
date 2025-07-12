package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Task2API/config"
	"Task2API/internal/handler"
	"Task2API/internal/middleware"
	"Task2API/internal/repository"
	"Task2API/internal/services"
	"Task2API/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.MustLoadConfig("config.yml")

	logger := &logger.StdLogger{}

	dbPool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		logger.Error("Unable to create connection pool: %v", err)
	}
	defer dbPool.Close()

	tokenRepo := repository.NewTokenRepository(dbPool, logger)

	targetAPIConfig := services.TargetAPIConfig{
		URL:           cfg.TargetAPI.URL,
		Authorization: cfg.TargetAPI.Authorization,
	}
	predictionService := services.NewPredictionService(targetAPIConfig, logger)

	h := handler.NewHandler(predictionService, logger)

	authMiddleware := middleware.NewAuthMiddleware(tokenRepo, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/predict/hba1c", h.PredictHBA1C)
	mux.HandleFunc("/predict/ldl", h.PredictLDL)
	mux.HandleFunc("/predict/ldll", h.PredictLDLL)
	mux.HandleFunc("/predict/ferr", h.PredictFERR)
	mux.HandleFunc("/predict/tg", h.PredictTG)
	mux.HandleFunc("/predict/hdl", h.PredictHDL)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      authMiddleware.Middleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Starting API gateway on :%s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server: %v", err)
		}
	}()

	<-done
	logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server gracefully: %v", err)
	}
	logger.Info("Server stopped")
}
