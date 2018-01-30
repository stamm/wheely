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
	"github.com/stamm/wheely/apis/distance/types"
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

func (svc Service) Name() string {
	return "Google"
}

func (svc Service) Calculate(ctx context.Context, start, end types.Point) (types.Result, error) {
	tr := &http.Transport{
		MaxIdleConns:    50,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
	}
	ret := types.Result{}
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%f,%f&destination=%f,%f&key=%s&units=metric", start.Lat, start.Long, end.Lat, end.Long, svc.token)
	log.Printf("do query GOOGLE url = %s\n", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ret, err
	}
	request = request.WithContext(ctx)
	resp, err := client.Do(request)
	if err != nil {
		return ret, err
	}

	var result response
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ret, err
	}

	if len(result.Routes) == 0 {
		return ret, errors.New("(routes) == 0")
	}
	if len(result.Routes[0].Legs) == 0 {
		return ret, errors.New("(legs) == 0")
	}
	return types.Result{
		Distance: result.Routes[0].Legs[0].Distance.Meters,
		Duration: result.Routes[0].Legs[0].Duration.Duration(),
	}, nil
}
