package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"Task2API/internal/interfaces"
	"Task2API/internal/models"
	"Task2API/pkg/logger"
)

type Handler struct {
	service interfaces.PredictionService
	logger  logger.Logger
}

func NewHandler(service interfaces.PredictionService, logger logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

//-----------------------------------------------------------------------------------------------

func (h *Handler) PredictHBA1C(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictHBA1C)
}

func (h *Handler) PredictLDL(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictLDL)
}

func (h *Handler) PredictLDLL(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictLDLL)
}

func (h *Handler) PredictFERR(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictFERR)
}

func (h *Handler) PredictTG(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictTG)
}

func (h *Handler) PredictHDL(w http.ResponseWriter, r *http.Request) {
	h.handlePrediction(w, r, h.service.PredictHDL)
}

//------------------------------------------------------------------------------------------

type predictionFunc func(context.Context, url.Values) ([]byte, error)

//------------------------------------------------------------------------------------------

func (h *Handler) handlePrediction(w http.ResponseWriter, r *http.Request, fn predictionFunc) {
	if r.Method != http.MethodGet {
		h.respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	response, err := fn(r.Context(), queryParams)
	if err != nil {
		h.logger.Error("Prediction error: %v", err)
		if strings.Contains(err.Error(), "service temporarily unavailable") {
			h.respondWithError(w, "Сервис прогнозирования временно недоступен. Пожалуйста, попробуйте позже.", http.StatusFailedDependency)
			return
		}
		h.respondWithError(w, fmt.Sprintf("Ошибка сервиса: %v", err), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: message})
}
