package fare

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kit/kit/endpoint"
	distancetypes "github.com/stamm/wheely/apis/distance/types"
	"github.com/stamm/wheely/apis/fare/types"
)

type FareService struct {
	distanceCalc endpoint.Endpoint
	tariff       types.Tariff
}

var (
	_ IFareService = FareService{}
)

func NewService(distanceCalc endpoint.Endpoint, tariff types.Tariff) FareService {
	return FareService{
		distanceCalc: distanceCalc,
		tariff:       tariff,
	}
}

func (svc FareService) Calculate(ctx context.Context, start, end distancetypes.Point) (types.Result, error) {
	result, err := svc.query(ctx, start, end)
	if err != nil {
		log.Printf("error from svc: %s", err)
		return types.Result{}, fmt.Errorf("Couldn't get result: %s", err.Error())
	}
	return types.Result{Amount: result}, nil
}

func (svc FareService) query(ctx context.Context, start, end distancetypes.Point) (int64, error) {
	request := distancetypes.CalculateRequest{
		StartLat:  start.Lat,
		StartLong: start.Long,
		EndLat:    end.Lat,
		EndLong:   end.Long,
	}

	resp, err := svc.distanceCalc(ctx, request)
	if err != nil {
		return 0, err
	}
	response := resp.(distancetypes.CalculateResponse)
	result := svc.tariff.Delivery + svc.tariff.ForMinute*(response.Duration/60) + svc.tariff.ForKm*(response.Distance/1000)
	if result < svc.tariff.Minimum {
		result = svc.tariff.Minimum
	}
	return result, nil
}
