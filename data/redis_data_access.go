package data

import "github.com/go-redis/redis/v8"

var (
	Rdb *redis.Client
)

func ConfigureCaching() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.201:30893",
		Username: "",
		Password: "",
		DB:       1, // use DB 1
	})
}
