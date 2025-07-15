package middleware

import (
	"net/http"
	"strconv"
	"time"

	"Task2API/pkg/logger"
	"go.uber.org/ratelimit"
)

type RateLimitMiddleware struct {
	limiter ratelimit.Limiter
	logger  logger.Logger
	rate    int
}

//------------------------------------------------------------------------------------------

func NewRateLimitMiddleware(rate int, logger logger.Logger) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		limiter: ratelimit.New(rate, ratelimit.Per(time.Second)),
		logger:  logger,
		rate:    rate,
	}
}

func (m *RateLimitMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		before := time.Now()
		m.limiter.Take()
		waitTime := time.Since(before)

		w.Header().Set("X-RateLimit-Limit", strconv.Itoa(m.rate))

		if waitTime > 0 {
			m.logger.Debug("Request delayed by %v due to rate limiting", waitTime)
		}

		next.ServeHTTP(w, r)
	})
}
