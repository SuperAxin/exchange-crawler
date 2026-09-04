package main

import (
	"context"
	"log"

	"binance-crawler/internal/binance"
	"binance-crawler/internal/crawler"
	"binance-crawler/internal/redis"
)

func main() {
	ctx := context.Background()

	binanceClient := binance.NewClient(
		"wss://stream.binance.com/stream",
	)
	defer binanceClient.Close()

	redisClient := redis.NewClient(
		"localhost:6379",
	)
	defer redisClient.Close()

	if err := binanceClient.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	err := binanceClient.Subscribe([]string{
		"!miniTicker@arr@3000ms",
		"btcusdt@aggTrade",
		"btcusdt@depth",
		"btcusdt@kline_1m",
	})

	if err != nil {
		log.Fatal(err)
	}

	app := crawler.NewCrawler(
		binanceClient,
		redisClient,
	)

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}