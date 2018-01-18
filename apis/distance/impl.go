package distance

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/stamm/wheely/apis/distance/services"
	"github.com/stamm/wheely/apis/distance/services/google"
	"github.com/stamm/wheely/apis/distance/services/osm"
	"github.com/stamm/wheely/apis/distance/types"
)

type DistanceService struct {
	services []services.Service
}

var (
	_ IDistanceService = DistanceService{}
)

const (
	timeout = 5 * time.Second
)

func NewService(googleToken, osmToken string) DistanceService {
	return DistanceService{
		services: []services.Service{
			google.New(googleToken),
			osm.New(osmToken),
		},
	}
}

func (svc DistanceService) Calculate(ctx context.Context, start, end types.Point) (types.Result, error) {
	for _, externalSvc := range svc.services {
		reqCtx, cancel := context.WithTimeout(ctx, timeout)
		result, err := externalSvc.Calculate(reqCtx, start, end)
		cancel()
		if err != nil {
			log.Printf("error from svc: %s", err)
			continue
		}
		return result, nil
	}
	return types.Result{}, errors.New("Couldn't get result")
}
