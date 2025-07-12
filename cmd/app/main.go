package main

import (
	"net/http"

	"Task2API/config"
	"Task2API/internal/handler"
	"Task2API/internal/middleware"
	"Task2API/internal/services"
	"Task2API/pkg/logger"
)

func main() {
	cfg := config.MustLoadConfig("config.yml")

	log := &logger.StdLogger{}

	targetAPIConfig := services.TargetAPIConfig{
		URL:           cfg.TargetAPI.URL,
		Authorization: cfg.TargetAPI.Authorization,
	}
	predictionService := services.NewPredictionService(targetAPIConfig, log)

	h := handler.NewHandler(predictionService, log)

	mux := http.NewServeMux()
	mux.HandleFunc("/predict/hba1c", h.PredictHBA1C)
	mux.HandleFunc("/predict/ldl", h.PredictLDL)
	mux.HandleFunc("/predict/ldll", h.PredictLDLL)
	mux.HandleFunc("/predict/ferr", h.PredictFERR)
	mux.HandleFunc("/predict/tg", h.PredictTG)
	mux.HandleFunc("/predict/hdl", h.PredictHDL)

	authMiddleware := middleware.AuthMiddleware(cfg.Auth.Token, log)
	handlerWithMiddleware := authMiddleware(mux)

	log.Info("Starting API gateway on :%s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, handlerWithMiddleware); err != nil {
		log.Error("Failed to start server: %v", err)
	}
}
