package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"Task2API/internal/repository"
	"Task2API/pkg/logger"
)

type AuthMiddleware struct {
	tokenRepo *repository.TokenRepository
	logger    logger.Logger
}

func NewAuthMiddleware(tokenRepo *repository.TokenRepository, logger logger.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		tokenRepo: tokenRepo,
		logger:    logger,
	}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			m.logger.Warn("Authorization header missing")
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			m.logger.Warn("Invalid authorization header format")
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		valid, err := m.tokenRepo.IsValidToken(ctx, token)
		if err != nil {
			m.logger.Error("Token validation error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !valid {
			m.logger.Warn("Invalid authorization token")
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
