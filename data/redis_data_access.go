package data

import "github.com/go-redis/redis/v8"

var (
	Rdb *redis.Client
)

func ConfigureCaching() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1, // use DB 1
	})
}
