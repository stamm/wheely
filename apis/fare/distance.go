package fare

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	distancetypes "github.com/stamm/wheely/apis/distance/types"
)

const (
	retry   = 3
	timeout = 2000 * time.Millisecond
)

func GetDistanceEndpoint(consulAddr string) endpoint.Endpoint {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	sdClient := consul.NewClient(client)
	instancer := consul.NewInstancer(sdClient, logger, "distance-api", nil, false)
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	return lb.Retry(retry, timeout, balancer)
}

func factory(instance string) (endpoint.Endpoint, io.Closer, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	u.Path = "/calculate"
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeResponse,
	).Endpoint(), nil, nil
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response distancetypes.CalculateResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
