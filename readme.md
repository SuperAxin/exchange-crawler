# Binance Crawler

使用 Go 開發的交易所即時行情爬蟲
透過 WebSocket 訂閱 Binance 即時市場資料，接收到行情後，使用 Redis Pub/Sub 將資料發布給其他服務使用。
