package services

import (
	"context"
	"time"
)

type Service interface {
	Calculate(ctx context.Context, startLat, startLong, endLat, endLong float64) (int64, time.Duration, error)
	// Name() String
}
