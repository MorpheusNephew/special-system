package redisclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// IClient an interface representing getting and setting data in Redis
type IClient interface {
	GetValue(key string)
	SetValue(key string)
}

// Client is the client for interacting with Redis
type Client struct {
	cache *redis.Client
}

// GetValue gets value from Redis with a given key
func (c *Client) GetValue(key string) {
	ctx := context.Background()

	defer ctx.Done()

	stuff := c.cache.Get(ctx, key)

	if stuff != nil {
		fmt.Printf("Retrieved from redis %v\n", stuff)
	} else {
		fmt.Println("Nothing retrieved from Redis")
	}

	fmt.Printf("GetValue was called with key: %v\n", key)
}

// SetValue sets value in Redis with a given key
func (c *Client) SetValue(key string) {
	ctx := context.Background()

	defer ctx.Done()

	fmt.Printf("SetValue was called with key: %v\n", key)
}
