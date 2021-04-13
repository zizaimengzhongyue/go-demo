package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const S = 1000000 * 1000

var ctx context.Context
var rdb *redis.Client

func init() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
