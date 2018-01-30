package stupid

import (
	"math"

	"github.com/stamm/wheely/apis/distance/cache"
	"github.com/stamm/wheely/apis/distance/types"
)

var (
	_ cache.ICache = &Cache{}
)

func New() *Cache {
	return &Cache{
		storage: make(map[types.Travel]types.Result),
	}
}

type Cache struct {
	storage map[types.Travel]types.Result
}

func (c *Cache) Set(k types.Travel, v types.Result) error {
	c.storage[k] = v
	return nil
}

func (c *Cache) Get(k types.Travel) (types.Result, error) {
	return c.storage[k], nil
}

func (c *Cache) Find(k types.Travel, precision float64) (types.Result, bool, error) {
	for travel, res := range c.storage {
		if (math.Abs(travel.Start.Lat-k.Start.Lat) <= precision) &&
			(math.Abs(travel.Start.Long-k.Start.Long) <= precision) &&
			(math.Abs(travel.End.Lat-k.End.Lat) <= precision) &&
			(math.Abs(travel.End.Long-k.End.Long) <= precision) {
			return res, true, nil
		}
	}
	return types.Result{}, false, nil
}
