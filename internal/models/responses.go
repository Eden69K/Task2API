package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type PredictionResponse struct {
	Result float64 `json:"result"`
}
