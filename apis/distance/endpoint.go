package distance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/stamm/wheely/apis/distance/types"
)

func MakeCalculationEndpoint(svc DistanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.CalculateRequest)
		result, err := svc.Calculate(ctx,
			types.NewPoint(req.StartLat, req.StartLong),
			types.NewPoint(req.EndLat, req.EndLong),
		)
		if err != nil {
			return types.CalculateResponse{
				Err: err.Error(),
			}, nil
		}
		return types.CalculateResponse{
			Distance: result.Distance,
			Duration: int64(result.Duration.Seconds()),
		}, nil
	}
}
