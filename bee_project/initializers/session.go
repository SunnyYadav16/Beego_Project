package initializers

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var redisConn *redis.Client

func SessionConnect() {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisConn.Ping(redisConn.Context()).Result()
	fmt.Println(pong, err)
}
