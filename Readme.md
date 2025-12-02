# 📦 Delta Engine — High-Performance Order Matching Engine (Go)

Delta Engine is a **lightweight, high-performance matching engine** written in Go.

It supports:

- **Limit orders (BUY/SELL)**
- **Price-time priority matching**
- **Max-heap bids & min-heap asks**
- **Partial fills**
- **Trade generation**
- **Complete order lifecycle management**
- **In-memory user + order storage**

This project simulates how real crypto & stock exchanges internally process orders.

---

## 🛠 Features

### ✔ Add Order
Validates and inserts a BUY or SELL order into the orderbook.

### ✔ Matching Engine
Matches incoming orders against existing orders based on price priority:

- BUY → matches against **best ask**
- SELL → matches against **best bid**
- Executes trades until:
  - one side is fully filled  
  - or prices no longer match  
  - or no liquidity available  

### ✔ Heap-Based Orderbook
- Bids → **max-heap**  
- Asks → **min-heap**  
Provides O(log n) insert/delete and O(1) best-price lookup.

### ✔ Partial Fills
If an order is partially filled, leftover quantity is automatically updated and reinserted.

### ✔ Order Lifecycle
- Insert  
- Match  
- Update  
- Remove  

### ✔ Simple In-Memory Storage
Stores:
- Users  
- Balances  
- Orders  

---

# 📂 Project Structure








## File Summary

- main.go : Entry point. Demonstrates order placement and matching.
- engine/orderbook.go : Core order book logic with Add/Cancel/Match.
- engine/match.go : Matching engine for bids and asks.
- engine/heap.go : Heap implementation for price-priority orders.
- engine/order.go : Order struct.
- engine/trade.go : Trade struct.
- utils/id.go : Generates unique IDs.
- storage/memory.go : In-memory storage for users/orders.
- engine/validate.go : Validates orders.

- 
