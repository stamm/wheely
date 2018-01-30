package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/stamm/wheely/apis/fare"
	"github.com/stamm/wheely/apis/fare/types"
)

func main() {
	consul := flag.String("consul", "consul:8500", "address for consul")
	flag.Parse()
	svc := fare.NewService(fare.GetDistanceEndpoint(*consul))

	calcHandler := httptransport.NewServer(
		fare.MakeCalculationEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)
	http.Handle("/calculate", calcHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
