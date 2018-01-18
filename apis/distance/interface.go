package distance

import (
	"context"

	"github.com/stamm/wheely/apis/distance/types"
)

type IDistanceService interface {
	Calculate(ctx context.Context, start, end types.Point) (types.Result, error)
}
