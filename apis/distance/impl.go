package distance

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/stamm/wheely/apis/distance/services"
	"github.com/stamm/wheely/apis/distance/services/google"
)

type DistanceService struct {
	services []services.Service
}

var (
	_ IDistanceService = DistanceService{}
)

func NewService(googleToken string) DistanceService {
	return DistanceService{
		services: []services.Service{
			google.New(googleToken),
		},
	}
}

func (svc DistanceService) Calculate(ctx context.Context, startLat, startLong, endLat, endLong float64) (int64, time.Duration, error) {
	for _, externalSvc := range svc.services {
		distance, duration, err := externalSvc.Calculate(ctx, startLat, startLong, endLat, endLong)
		if err != nil {
			log.Printf("error from svc: %s", err)
			continue
		}
		return distance, duration, nil
	}
	return 0, 0, errors.New("Couldn't get result")
}
