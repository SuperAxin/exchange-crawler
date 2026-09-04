package binance

import "github.com/gorilla/websocket"

type Client struct {
	baseURL string
	conn    *websocket.Conn
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
	}
}

func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}

	return c.conn.Close()
}