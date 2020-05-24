package redisclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
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
