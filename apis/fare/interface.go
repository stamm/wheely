package fare

import (
	"context"

	"github.com/stamm/wheely/apis/distance/types"
	faretypes "github.com/stamm/wheely/apis/fare/types"
)

type IFareService interface {
	Calculate(ctx context.Context, start, end types.Point) (faretypes.Result, error)
}
