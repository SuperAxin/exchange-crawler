package binance

import (
	"context"

	"github.com/gorilla/websocket"
)

func (c *Client) Connect(ctx context.Context) error {
	conn, _, err := websocket.DefaultDialer.DialContext(
		ctx,
		c.baseURL,
		nil,
	)

	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

func (c *Client) Subscribe(streams []string) error {
	request := SubscribeRequest{
		Method: "SUBSCRIBE",
		Params: streams,
		ID:     1,
	}

	return c.conn.WriteJSON(request)
}

func (c *Client) Read(
	ctx context.Context,
	messages chan<- []byte,
) error {

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}

		select {
		case messages <- message:

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}