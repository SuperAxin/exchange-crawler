package crawler

import (
	"context"

	"binance-crawler/internal/binance"
	"binance-crawler/internal/redis"
)

type Crawler struct {
	binance *binance.Client
	redis   *redis.Client
}

func NewCrawler(
	binanceClient *binance.Client,
	redisClient *redis.Client,
) *Crawler {

	return &Crawler{
		binance: binanceClient,
		redis:   redisClient,
	}
}

func (c *Crawler) Run(ctx context.Context) error {
	messages := make(chan []byte, 100)

	go func() {
		if err := c.binance.Read(ctx, messages); err != nil {
			return
		}
	}()

	for {
		select {

		case message := <-messages:

			if err := c.redis.Publish(
				ctx,
				"binance:trade",
				message,
			); err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}