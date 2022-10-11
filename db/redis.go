package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewClient() *redis.Client { // 實體化redis.Client 並返回實體的位址
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
