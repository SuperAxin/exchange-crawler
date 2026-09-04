package redis

import (
	"context"
)

func (c *Client) Publish(
	ctx context.Context,
	channel string,
	message []byte,
) error {

	return c.client.Publish(
		ctx,
		channel,
		message,
	).Err()
}