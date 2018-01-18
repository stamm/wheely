package services

import (
	"context"

	"github.com/stamm/wheely/apis/distance/types"
)

type Service interface {
	Calculate(ctx context.Context, start, end types.Point) (types.Result, error)
	Name() string
}
