package repository

import (
	"context"
	//"time"

	"Task2API/pkg/logger"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	pool   *pgxpool.Pool
	logger logger.Logger
}

func NewTokenRepository(pool *pgxpool.Pool, logger logger.Logger) *TokenRepository {
	return &TokenRepository{
		pool:   pool,
		logger: logger,
	}
}

func (r *TokenRepository) IsValidToken(ctx context.Context, token string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 FROM public.auth_tokens 
		WHERE token = $1 AND expires_at > NOW()
	)`

	err := r.pool.QueryRow(ctx, query, token).Scan(&exists)
	if err != nil {
		r.logger.Error("Failed to check token validity: %v", err)
		return false, err
	}

	return exists, nil
}

func (r *TokenRepository) Close() {
	r.pool.Close()
}
