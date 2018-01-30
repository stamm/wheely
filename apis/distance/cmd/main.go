package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"github.com/stamm/wheely/apis/distance"
	"github.com/stamm/wheely/apis/distance/cache/stupid"
	"github.com/stamm/wheely/apis/distance/types"
)

func main() {
	var svc distance.IDistanceService
	svc = distance.NewService(
		os.Getenv("GOOGLE_TOKEN"),
		os.Getenv("OSM_TOKEN"),
	)
	// enable cache with precision 0.002
	svc = distance.NewCacheMiddleware(stupid.New(), 0.002, svc)

	calculateHandler := httptransport.NewServer(
		distance.MakeCalculationEndpoint(svc),
		decodeCalculatorRequest,
		encodeResponse,
	)

	http.Handle("/calculate", calculateHandler)
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = "consul:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	sdClient := consul.NewClient(client)
	registar := consul.NewRegistrar(sdClient, &api.AgentServiceRegistration{
		ID:      "1",
		Name:    "distance-api",
		Port:    8080,
		Address: hostname,
	}, kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))
	registar.Register()
	defer registar.Deregister()

	log.Printf("hostname: %s, envs: %v\n", hostname, os.Environ())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeCalculatorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
