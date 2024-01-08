package db

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectDB() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println("Connected to Redis")
}
