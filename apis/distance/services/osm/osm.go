package osm

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
	return "OSM"
}

func (svc Service) Calculate(ctx context.Context, start, end types.Point) (types.Result, error) {
	tr := &http.Transport{
		MaxIdleConns:    50,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
	}
	url := fmt.Sprintf("https://api.openrouteservice.org/directions?api_key=%s&coordinates=%f,%f|%f,%f&profile=driving-car&preference=fastest&units=m&language=en&geometry=false", svc.token, start.Long, start.Lat, end.Long, end.Lat)
	log.Printf("do query OSM url = %s\n", url)
	resp, err := client.Get(url)
	ret := types.Result{}
	if err != nil {
		return ret, err
	}

	var result response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ret, err
	}
	if len(result.Routes) == 0 {
		return ret, errors.New("(routes) == 0")
	}
	return types.Result{
		Distance: int64(result.Routes[0].Summary.Distance),
		Duration: time.Duration(result.Routes[0].Summary.Duration) * time.Second,
	}, nil
}
