package redis

import "fmt"

// IClient an interface representing getting and setting data in Redis
type IClient interface {
	GetValue(key string)
	SetValue(key string)
}

// Client is the client for interacting with Redis
type Client struct {
}

// GetValue gets value from Redis with a given key
func (c *Client) GetValue(key string) {
	fmt.Printf("GetValue was called with key: %v\n", key)
}

// SetValue sets value in Redis with a given key
func (c *Client) SetValue(key string) {
	fmt.Printf("SetValue was called with key: %v\n", key)
}
