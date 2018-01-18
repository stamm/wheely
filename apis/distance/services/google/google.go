package google

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stamm/wheely/apis/distance/services"
)

type Service struct {
	token string
}

var (
	_ services.Service = Service{}
)

func New(token string) Service {
	return Service{
		token: token,
	}
}

func (svc Service) Calculate(ctx context.Context, startLat, startLong, endLat, endLong float64) (int64, time.Duration, error) {
	tr := &http.Transport{
		MaxIdleConns:    50,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
	}
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%f,%f&destination=%f,%f&key=%s&units=metric", startLat, startLong, endLat, endLong, svc.token)
	log.Printf("url = %+v\n", url)
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, err
	}

	var result response
	json.NewDecoder(resp.Body).Decode(&result)
	if len(result.Routes) == 0 {
		return 0, 0, errors.New("(routes) == 0")
	}
	if len(result.Routes[0].Legs) == 0 {
		return 0, 0, errors.New("(legs) == 0")
	}
	return result.Routes[0].Legs[0].Distance.Meters, result.Routes[0].Legs[0].Duration.Duration(), nil
}
