package distance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeCalculationEndpoint(svc DistanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		distance, duration, err := svc.Calculate(ctx, req.StartLat, req.StartLong, req.EndLat, req.EndLong)
		if err != nil {
			return CalculateResponse{
				Err: err.Error(),
			}, nil
		}
		return CalculateResponse{
			Distance: distance,
			Duration: int64(duration.Seconds()),
		}, nil
	}
}
