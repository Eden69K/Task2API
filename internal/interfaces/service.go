package interfaces

import (
	"context"
	"net/url"
)

type PredictionService interface {
	PredictHBA1C(ctx context.Context, params url.Values) ([]byte, error)
	PredictLDL(ctx context.Context, params url.Values) ([]byte, error)
	PredictLDLL(ctx context.Context, params url.Values) ([]byte, error)
	PredictFERR(ctx context.Context, params url.Values) ([]byte, error)
	PredictTG(ctx context.Context, params url.Values) ([]byte, error)
	PredictHDL(ctx context.Context, params url.Values) ([]byte, error)
}
