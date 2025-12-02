# 📦 Delta Engine — High-Performance Order Matching Engine (Go)

Delta Engine is a lightweight, high-performance order-matching engine written in Go. It simulates how crypto and stock exchanges handle limit orders internally using price-time priority. It supports BUY and SELL orders, automatic matching, partial fills, trade generation, and in-memory order/user storage.

---

## 🛠 Features

- Limit BUY/SELL orders
- Automatic matching engine
- Price-time priority
- Partial fills
- Bids = Max-Heap, Asks = Min-Heap
- Trade generation
- Order and user storage
- Clean modular architecture

---

## 🚀 Running the Engine

1. Clone the repository:
git clone https://github.com/ParthAggarwal16/Delta-Engine
cd Delta-Engine
go run main.go

## Example Trade
Trade executed:
BUY  -> O-12345
SELL -> O-67890
Price:    100.0
Quantity:  2

Best Bid: 99.0 (qty 3)
Best Ask: 105.0 (qty 5)

Future Improvements
	•	REST API (Gin/Fiber)
	•	WebSocket live orderbook feed
	•	Persistent storage (PostgreSQL/SQLite)
	•	Market orders
	•	Fee calculation
	•	Historical trades
	•	Metrics dashboard
	•	Docker + Render deployment
	•	CLI trading interface
