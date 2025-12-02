package engine

import "time"

// Exported constants for buy/sell
const (
	BUY  = "buy"
	SELL = "sell"
)

// OrderBook holds bids, asks and order map
type OrderBook struct {
	Bids   *OrderHeap
	Asks   *OrderHeap
	Orders map[string]*Order
}

// AddOrder validates, matches, and stores leftover
func (ob *OrderBook) AddOrder(o *Order) []*Trade {
	trades := ob.Match(o) // match incoming order

	if o.Quantity > 0 { // store leftover
		ob.Orders[o.ID] = o
		if o.Side == BUY {
			ob.Bids.Insert(o)
		} else {
			ob.Asks.Insert(o)
		}
	}

	return trades
}

// CancelOrder removes order by ID
func (ob *OrderBook) CancelOrder(id string) {
	o := ob.Orders[id] // defining the ID of order
	if o == nil {
		return
	}
	if o.Side == BUY {
		ob.Bids.Remove(o)
	} else {
		ob.Asks.Remove(o)
	}
	delete(ob.Orders, id)
}

// GetBestBid returns top bid
func (ob *OrderBook) GetBestBid() *Order {
	return ob.Bids.PeekRoot()
}

// GetBestAsk returns top ask
func (ob *OrderBook) GetBestAsk() *Order {
	return ob.Asks.PeekRoot()
}

// Match performs matching and returns trades
func (ob *OrderBook) Match(o *Order) []*Trade {
	var trades []*Trade //defining trades

	if o.Side == BUY { //if incoming buy

		for o.Quantity > 0 {

			bestAsk := ob.GetBestAsk() //calling getbestask inside this function

			if bestAsk == nil || o.Price < bestAsk.Price {
				break
			}
			tradeQty := min(o.Quantity, bestAsk.Quantity) //quantitty to be traded will be minimum of order quantity and bestAsk quantity

			trade := &Trade{ //defining trade instance
				BuyOrderID:  o.ID,
				SellOrderID: bestAsk.ID,
				Price:       bestAsk.Price,
				Quantity:    tradeQty,
				Timestamp:   time.Now().Unix(),
			}

			trades = append(trades, trade) // appending instance to trades

			o.Quantity -= tradeQty // updatign quantities
			bestAsk.Quantity -= tradeQty

			if bestAsk.Quantity == 0 { //remove bestAsk after the quantity becomes zero
				ob.Asks.Remove(bestAsk)
				delete(ob.Orders, bestAsk.ID)
			} else { //if it doesnt go down to zero , update the bestAsk
				ob.Asks.Update(bestAsk)
			}
		}
	} else { // SELL

		for o.Quantity > 0 {

			bestBid := ob.GetBestBid() // calling GetBestBid inside the function

			if bestBid == nil || o.Price > bestBid.Price { // if no bid exists or order price is greater than bestBid price then return empty trade
				break
			}
			tradeQty := min(o.Quantity, bestBid.Quantity) //quantitty to be traded will be minimum of order quantity and besbid quantity

			trade := &Trade{ //defining a trade instance
				BuyOrderID:  bestBid.ID,
				SellOrderID: o.ID,
				Price:       bestBid.Price,
				Quantity:    tradeQty,
				Timestamp:   time.Now().Unix(),
			}

			trades = append(trades, trade) // adding the instance to the trades

			o.Quantity -= tradeQty // updating quantity after trade
			bestBid.Quantity -= tradeQty

			if bestBid.Quantity == 0 { //remove bestbid after the quantity becomes zero
				ob.Bids.Remove(bestBid)
				delete(ob.Orders, bestBid.ID)
			} else { //if not then update bestbid
				ob.Bids.Update(bestBid)
			}
		}
	}

	return trades
}

// helper min function
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
