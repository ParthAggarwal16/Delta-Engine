package engine

// Trade represents an executed trade between a buy order and a sell order
type Trade struct {
	BuyOrderID  string  // ID of the buy order
	SellOrderID string  // ID of the sell order
	Price       float64 // execution price
	Quantity    float64 // quantity executed in this trade
	Timestamp   int64   // epoch time when trade occurred
}
