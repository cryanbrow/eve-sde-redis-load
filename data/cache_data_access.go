package data

import (
	"github.com/patrickmn/go-cache"
)

var (
	NonExpiringCache *cache.Cache
	TypeCache        *cache.Cache
	ExpiringCache    *cache.Cache
)

func ConfigureInMemoryCaching() {
	NonExpiringCache = cache.New(cache.NoExpiration, cache.NoExpiration)
	TypeCache = cache.New(cache.DefaultExpiration, cache.NoExpiration)
	ExpiringCache = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
}
