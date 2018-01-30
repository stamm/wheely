package distance

import (
	"context"
	"log"

	"github.com/stamm/wheely/apis/distance/cache"
	"github.com/stamm/wheely/apis/distance/types"
)

type CacheMiddleware struct {
	cache     cache.ICache
	precision float64
	next      IDistanceService
}

func (mw CacheMiddleware) Calculate(ctx context.Context, start, end types.Point) (types.Result, error) {
	res, ok, err := mw.cache.Find(types.Travel{Start: start, End: end}, mw.precision)
	if err == nil && ok {
		log.Printf("Found from cache from %v to %v\n", start, end)
		return res, nil
	}
	result, err := mw.next.Calculate(ctx, start, end)
	if err == nil {
		errCache := mw.cache.Set(types.Travel{Start: start, End: end}, result)
		if errCache != nil {
			log.Printf("Error in set cache: %s\n", errCache.Error())
		}
	}
	return result, err
}

func NewCacheMiddleware(cache cache.ICache, precision float64, next IDistanceService) CacheMiddleware {
	return CacheMiddleware{
		cache:     cache,
		precision: precision,
		next:      next,
	}
}
