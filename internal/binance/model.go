package binance

type StreamMessage struct {
	Stream string        `json:"stream"`
	Data   AggTradeEvent `json:"data"`
}

type AggTradeEvent struct {
	EventType      string `json:"e"`
	EventTime      int64  `json:"E"`
	Symbol         string `json:"s"`
	AggTradeID     int64  `json:"a"`
	Price          string `json:"p"`
	Quantity       string `json:"q"`
	FirstTradeID   int64  `json:"f"`
	LastTradeID    int64  `json:"l"`
	TradeTime      int64  `json:"T"`
	BuyerIsMaker   bool   `json:"m"`
	BestPriceMatch bool   `json:"M"`
}

type SubscribeRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}