package redisclient

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/morpheusnephew/qotd/internal/variables"
)

// IClient an interface representing getting and setting data in Redis
type IClient interface {
	GetValue(key string) ([]byte, error)
	SetValue(key string, value []byte, t *time.Duration) (string, error)
	GetInitialized() bool
	SetInitialized(value bool)
}

// Client is the client for interacting with Redis
type Client struct {
	cache       *redis.Client
	initialized bool
}

// IClientFactory is an interface for client factory
type IClientFactory interface {
	GetRedisClient() IClient
}

// ClientFactory is a struct used to get a redis client
type ClientFactory struct {
	client IClient
}

// GetRedisClient is a method to get the Redis client
func (cf *ClientFactory) GetRedisClient() IClient {
	if cf.client == nil || !cf.client.GetInitialized() {
		endpoint := fmt.Sprintf("%v:%v", variables.RedisEndpoint, variables.RedisPort)
		cache := redis.NewClient(&redis.Options{
			Addr: endpoint,
		})

		ctx := context.Background()
		defer ctx.Done()

		_, err := cache.Ping(ctx).Result()

		cf.client = &Client{
			cache: cache,
		}

		if err != nil {
			cf.client.SetInitialized(false)
		} else {
			cf.client.SetInitialized(true)
		}
	}

	return cf.client
}

// GetValue gets value from Redis with a given key
func (c *Client) GetValue(key string) ([]byte, error) {
	if !c.initialized {
		return []byte{}, nil
	}

	ctx := context.Background()

	defer ctx.Done()

	dataValue, err := c.cache.Get(ctx, key).Result()

	return []byte(dataValue), err
}

// SetValue sets value in Redis with a given key
func (c *Client) SetValue(key string, value []byte, t *time.Duration) (string, error) {
	if !c.initialized {
		return "", nil
	}

	ctx := context.Background()

	defer ctx.Done()

	if t == nil {
		*t = time.Hour * 24
	}

	return c.cache.Set(ctx, key, value, *t).Result()
}

// GetInitialized gets the initialized value
func (c *Client) GetInitialized() bool {
	return c.initialized
}

// SetInitialized sets the initialized value
func (c *Client) SetInitialized(value bool) {
	c.initialized = value
}
