package fare

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	distancetypes "github.com/stamm/wheely/apis/distance/types"
	"github.com/stamm/wheely/apis/fare/types"
)

func MakeCalculationEndpoint(svc IFareService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.CalculateRequest)
		result, err := svc.Calculate(ctx,
			distancetypes.NewPoint(req.StartLat, req.StartLong),
			distancetypes.NewPoint(req.EndLat, req.EndLong),
		)
		if err != nil {
			return types.CalculateResponse{
				Err: err.Error(),
			}, nil
		}
		return result, nil
	}
}
