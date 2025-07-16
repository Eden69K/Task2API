package repository

import (
	"context"
	"fmt"

	"Task2API/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	pool    *pgxpool.Pool
	logger  logger.Logger
	appName string
}

//------------------------------------------------------------------------------------------

func NewTokenRepository(pool *pgxpool.Pool, appName string, logger logger.Logger) *TokenRepository {
	config := pool.Config()
	config.ConnConfig.RuntimeParams["application_name"] = appName

	return &TokenRepository{
		pool:    pool,
		logger:  logger,
		appName: appName,
	}
}

func (r *TokenRepository) IsValidToken(ctx context.Context, token string) (bool, error) {
	r.logger.Debug("Checking token for app: %s", r.appName)

	query := `SELECT EXISTS(
		SELECT 1 FROM public.auth_tokens 
		WHERE token = $1 AND expires_at > NOW()
	)`

	var exists bool
	err := r.pool.QueryRow(ctx, query, token).Scan(&exists)
	if err != nil {
		r.logger.Error("DB error [app:%s]: %v", r.appName, err)
		return false, fmt.Errorf("DB error [%s]: %w", r.appName, err)
	}

	return exists, nil
}

func (r *TokenRepository) Close() {
	r.pool.Close()
}
