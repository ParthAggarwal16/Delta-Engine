package engine

// matchBid matches an incoming buy order with existing asks
func (ob *OrderBook) matchBid(o *Order) []*Trade {
	trades := []*Trade{}

	for {
		bestAsk := ob.GetBestAsk() //call GetBestAsk from orderbook

		if bestAsk == nil { // if no bestasks present , return empty trades
			break
		}

		if o.Price < bestAsk.Price { // if bestask is more than the price , return the empty trade
			break
		}

		tradeQty := min(o.Quantity, bestAsk.Quantity) //the minimum quantity between the bestask and the order will be traded

		t := &Trade{ //defining the trade instance
			BuyOrderID:  o.ID,
			SellOrderID: bestAsk.ID,
			Price:       bestAsk.Price,
			Quantity:    tradeQty,
		}

		trades = append(trades, t) //appending the successful trade

		o.Quantity -= tradeQty //updating the quantities
		bestAsk.Quantity -= tradeQty

		if bestAsk.Quantity == 0 { //if the ask quantity becomes zero , remove the ask
			ob.Asks.Remove(bestAsk)
		} else {
			ob.Asks.Update(bestAsk) // if not , update the bestask
		}

		if o.Quantity == 0 { //if the order quantity is zero , retiurn the empty trade
			break
		}
	}
	return trades
}

// matchask matches an incoming sell order with existing bids
func (ob *OrderBook) matchask(o *Order) []*Trade {
	trades := []*Trade{}

	for {
		bestBid := ob.GetBestBid() //call GetBestBid from orderbook

		if bestBid == nil { // if no bestbids present , return empty trades
			break
		}

		if bestBid.Price < o.Price { // if bestbid is less than the price , return the empty trade
			break
		}

		tradeQty := min(bestBid.Quantity, o.Quantity) //the minimum quantity between the bestbid and the order will be traded

		t := &Trade{ //defining the trade instance
			BuyOrderID:  bestBid.ID,
			SellOrderID: o.ID,
			Price:       bestBid.Price,
			Quantity:    tradeQty,
		}

		trades = append(trades, t) //appending the successful trade

		bestBid.Quantity -= tradeQty //updating the quantites
		o.Quantity -= tradeQty

		if bestBid.Quantity == 0 { //if the bid quantity is zero , retiurn the empty trade
			ob.Bids.Remove(bestBid)
		} else { //else just update the bestbid
			ob.Bids.Update(bestBid)
		}

		if o.Quantity == 0 { // if order quantity is zero , return empty trade
			break
		}
	}
	return trades
}

func generateTrade(buy, sell *Order, qty float64) *Trade { //function to generate trade

	return &Trade{
		BuyOrderID:  buy.ID,
		SellOrderID: sell.ID,
		Price:       sell.Price, // or buy.Price depending on convention
		Quantity:    qty,
	}
}
