package cache

import "github.com/stamm/wheely/apis/distance/types"

type ICache interface {
	Set(types.Travel, types.Result) error
	Get(types.Travel) (types.Result, error)
	Find(types.Travel, float64) (types.Result, bool, error)
}
