package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/stamm/wheely/apis/distance"
	"github.com/stamm/wheely/apis/distance/types"
)

func main() {
	svc := distance.NewService(
		os.Getenv("GOOGLE_TOKEN"),
		os.Getenv("OSM_TOKEN"),
	)

	calculateHandler := httptransport.NewServer(
		distance.MakeCalculationEndpoint(svc),
		decodeCalculatorRequest,
		encodeResponse,
	)

	http.Handle("/calculate", calculateHandler)
	log.Println("msg", "HTTP", "addr", ":8080")
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
