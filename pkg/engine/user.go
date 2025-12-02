package engine

// User represents a participant in the trading system
type User struct {
	ID      string            // Unique identifier for the user
	Name    string            // User's name (optional, can be used for UI)
	Orders  map[string]*Order // Map of order IDs to Order objects placed by the user
	Balance float64           // Optional: available balance for trading
}

// NewUser creates and returns a new user with an initialized orders map
func NewUser(id string, name string, balance float64) *User {
	return &User{
		ID:      id,
		Name:    name,
		Orders:  make(map[string]*Order),
		Balance: balance,
	}
}

// AddOrder adds an order to the user's order map
func (u *User) AddOrder(o *Order) {
	u.Orders[o.ID] = o
}

// RemoveOrder deletes an order from the user's order map
func (u *User) RemoveOrder(orderID string) {
	delete(u.Orders, orderID)
}
