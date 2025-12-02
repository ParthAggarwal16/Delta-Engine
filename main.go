package main

import (
	"fmt"
	"time"

	"DeltaEngine/pkg/engine"
)

func main() {
	// create order book
	orderBook := &engine.OrderBook{
		Bids:   engine.NewOrderHeap(engine.BidLess),
		Asks:   engine.NewOrderHeap(engine.AskLess),
		Orders: make(map[string]*engine.Order),
	}

	// create some orders
	order1 := &engine.Order{
		ID:        "o1",
		Price:     100,
		Quantity:  10,
		Side:      engine.BUY,
		Timestamp: time.Now().Unix(),
	}

	order2 := &engine.Order{
		ID:        "o2",
		Price:     95,
		Quantity:  5,
		Side:      engine.SELL,
		Timestamp: time.Now().Unix(),
	}

	// add orders
	trades := orderBook.AddOrder(order1)
	trades = append(trades, orderBook.AddOrder(order2)...)

	// print trades
	for _, t := range trades {
		fmt.Printf("Trade: BuyOrder %s, SellOrder %s, Price %.2f, Quantity %.2f\n",
			t.BuyOrderID, t.SellOrderID, t.Price, t.Quantity)
	}

	// print best bid and ask
	bestBid := orderBook.GetBestBid()
	bestAsk := orderBook.GetBestAsk()

	if bestBid != nil {
		fmt.Printf("Best Bid: %.2f\n", bestBid.Price)
	} else {
		fmt.Println("No bids")
	}

	if bestAsk != nil {
		fmt.Printf("Best Ask: %.2f\n", bestAsk.Price)
	} else {
		fmt.Println("No asks")
	}
}
