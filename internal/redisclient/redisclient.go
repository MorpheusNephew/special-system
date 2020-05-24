package redisclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/morpheusnephew/qotd/internal/utils"
	"github.com/morpheusnephew/qotd/internal/variables"
)

func init() {
	endpoint := fmt.Sprintf("%v:%v", variables.RedisEndpoint, variables.RedisPort)
	cache := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	ctx := context.Background()
	defer ctx.Done()

	_, err := cache.Ping(ctx).Result()

	utils.PanicIfError(err)

	client = &Client{
		cache: cache,
	}
}

var client IClient

// GetRedisClient returns the redis client
func GetRedisClient() IClient {
	return client
}
