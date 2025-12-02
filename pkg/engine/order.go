package engine

// Order represents a single buy or sell order in the order book

type Order struct {
	ID        string  // unique order ID
	Price     float64 // price per unit
	Quantity  float64 // remaining quantity to be filled
	Side      string  // BUY or SELL
	Timestamp int64   // epoch time when order was created
	Index     int     // index in the heap (used for heap operations)
}
