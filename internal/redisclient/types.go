package redisclient

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/morpheusnephew/qotd/internal/utils"
	"github.com/morpheusnephew/qotd/internal/variables"
)

// IClient an interface representing getting and setting data in Redis
type IClient interface {
	GetValue(key string) ([]byte, error)
	SetValue(key string, value []byte) (string, error)
}

// Client is the client for interacting with Redis
type Client struct {
	cache *redis.Client
}

// IClientFactory is an interface for client factory
type IClientFactory interface {
	GetRedisClient() IClient
}

// ClientFactory is a struct used to get a redis client
type ClientFactory struct {
	client IClient
}

// GetValue gets value from Redis with a given key
func (c *Client) GetValue(key string) ([]byte, error) {
	ctx := context.Background()

	defer ctx.Done()

	dataValue, err := c.cache.Get(ctx, key).Result()

	return []byte(dataValue), err
}

// SetValue sets value in Redis with a given key
func (c *Client) SetValue(key string, value []byte) (string, error) {
	ctx := context.Background()

	defer ctx.Done()

	return c.cache.Set(ctx, key, value, time.Hour*2).Result()
}

// GetRedisClient is a method to get the Redis client
func (cf *ClientFactory) GetRedisClient() IClient {
	if cf.client == nil {
		endpoint := fmt.Sprintf("%v:%v", variables.RedisEndpoint, variables.RedisPort)
		cache := redis.NewClient(&redis.Options{
			Addr: endpoint,
		})

		ctx := context.Background()
		defer ctx.Done()

		_, err := cache.Ping(ctx).Result()

		utils.PanicIfError(err)

		cf.client = &Client{
			cache: cache,
		}
	}

	return cf.client
}
