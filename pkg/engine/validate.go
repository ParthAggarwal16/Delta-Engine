package engine

import (
	"fmt"
	"math"
)

func (ob *OrderBook) validate(o *Order) error {
	if o == nil { //if there is no order
		return fmt.Errorf("order is nill")
	}
	if o.Side != "buy" && o.Side != "sell" { //checks if the side is valid
		return fmt.Errorf("the order is invalid")
	}
	if o.Price <= 0 || math.IsNaN(o.Price) || math.IsInf(o.Price, 0) { //checks if the price is valid
		return fmt.Errorf("the price is invalid")
	}
	if o.Quantity <= 0 || math.IsNaN(o.Quantity) || math.IsInf(o.Quantity, 0) { // check if the quantity is valid
		return fmt.Errorf("the quantity is invalid")
	}
	if o.ID == "" {
		return fmt.Errorf("the order ID cannot be emoty")
	}
	_, exists := ob.Orders[o.ID] // checking if it already exists using a placeholder becuase we dont care about the actual vakue
	if exists {
		return fmt.Errorf("the order ID already exists")
	}
	return nil
}
