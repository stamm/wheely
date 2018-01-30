package fare

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/stamm/wheely/apis/distance/types"
	faretypes "github.com/stamm/wheely/apis/fare/types"
)

type FareService struct {
	distanceCalc endpoint.Endpoint
}

var (
	_ IFareService = FareService{}
)

func NewService(distanceCalc endpoint.Endpoint) FareService {
	return FareService{
		distanceCalc: distanceCalc,
	}
}

func (svc FareService) Calculate(ctx context.Context, start, end types.Point) (faretypes.Result, error) {
	result, err := svc.query(ctx, start, end)
	if err != nil {
		log.Printf("error from svc: %s", err)
		return faretypes.Result{}, fmt.Errorf("Couldn't get result: %s", err.Error())
	}
	return faretypes.Result{Amount: result}, nil
}

func (svc FareService) query(ctx context.Context, start, end types.Point) (int64, error) {
	request := types.CalculateRequest{
		StartLat:  start.Lat,
		StartLong: start.Long,
		EndLat:    end.Lat,
		EndLong:   end.Long,
	}

	resp, err := svc.distanceCalc(ctx, request)
	if err != nil {
		return 0, err
	}
	response := resp.(types.CalculateResponse)
	result := 150 + 15*(response.Duration/60) + 38*(response.Distance/1000)
	if result < 299 {
		result = 299
	}
	return result, nil
}
