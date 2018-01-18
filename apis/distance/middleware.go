package distance

import (
	"context"
	"log"
	"math"

	"github.com/stamm/wheely/apis/distance/types"
)

type CacheMiddleware struct {
	cache     map[types.Travel]types.Result
	precision float64
	next      IDistanceService
}

func (mw CacheMiddleware) Calculate(ctx context.Context, start, end types.Point) (types.Result, error) {
	for travel, res := range mw.cache {
		if (math.Abs(travel.Start.Lat-start.Lat) <= mw.precision) &&
			(math.Abs(travel.Start.Long-start.Long) <= mw.precision) &&
			(math.Abs(travel.End.Lat-end.Lat) <= mw.precision) &&
			(math.Abs(travel.End.Long-end.Long) <= mw.precision) {
			log.Println("Found from cache")
			return res, nil
		}
	}
	result, err := mw.next.Calculate(ctx, start, end)
	if err == nil {
		mw.cache[types.Travel{Start: start, End: end}] = result
	}
	return result, err
}

func NewCacheMiddleware(cache map[types.Travel]types.Result, precision float64, next IDistanceService) CacheMiddleware {
	return CacheMiddleware{
		cache:     cache,
		precision: precision,
		next:      next,
	}
}
