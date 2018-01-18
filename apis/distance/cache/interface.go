package cache

import "github.com/stamm/wheely/apis/distance/types"

type Cache interface {
	Set(types.Travel, types.Result) error
	Get(types.Travel) (types.Result, error)
}
