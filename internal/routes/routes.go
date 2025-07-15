package routes

import (
	"net/http"

	"Task2API/internal/handler"
	"Task2API/internal/middleware"
)

type Router struct {
	handler        *handler.Handler
	authMiddleware *middleware.AuthMiddleware
	rateLimiter    *middleware.RateLimitMiddleware
}

//------------------------------------------------------------------------------------------

func NewRouter(h *handler.Handler, auth *middleware.AuthMiddleware, rl *middleware.RateLimitMiddleware) *Router {
	return &Router{
		handler:        h,
		authMiddleware: auth,
		rateLimiter:    rl,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/predict/hba1c", r.handler.PredictHBA1C)
	mux.HandleFunc("/predict/ldl", r.handler.PredictLDL)
	mux.HandleFunc("/predict/ldll", r.handler.PredictLDLL)
	mux.HandleFunc("/predict/ferr", r.handler.PredictFERR)
	mux.HandleFunc("/predict/tg", r.handler.PredictTG)
	mux.HandleFunc("/predict/hdl", r.handler.PredictHDL)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	handlerChain := r.rateLimiter.Middleware(
		r.authMiddleware.Middleware(mux),
	)

	return handlerChain
}
