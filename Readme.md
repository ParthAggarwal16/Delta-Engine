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
