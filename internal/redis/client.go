package redis

import (
	goredis "github.com/redis/go-redis/v9"
)

type Client struct {
	client *goredis.Client
}

func NewClient(addr string) *Client {
	return &Client{
		client: goredis.NewClient(&goredis.Options{
			Addr: addr,
		}),
	}
}

func (c *Client) Close() error {
	return c.client.Close()
}