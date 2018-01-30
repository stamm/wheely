package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/stamm/wheely/apis/fare"
	"github.com/stamm/wheely/apis/fare/types"
)

func main() {
	consul := flag.String("consul", "consul:8500", "address for consul")
	tariffFile := flag.String("tariff", "/cfg/tariff.json", "file to tarif")
	flag.Parse()
	tariff, err := getTariff(*tariffFile)
	if err != nil {
		panic(err)
	}

	svc := fare.NewService(fare.GetDistanceEndpoint(*consul), tariff)

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

func getTariff(file string) (types.Tariff, error) {
	var tariff types.Tariff
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return types.Tariff{}, err
	}
	err = json.Unmarshal(raw, &tariff)
	if err != nil {
		return types.Tariff{}, err
	}
	return tariff, nil
}
