package distance

import (
	"context"
	"time"
)

type IDistanceService interface {
	Calculate(ctx context.Context, startLat, startLong, endLat, endLong float64) (int64, time.Duration, error)
}
