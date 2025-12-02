package storage

import (
	"DeltaEngine/pkg/engine"
	"sync"
)

// In-memory storage for users and orders
type Memory struct {
	Users  map[string]*engine.User
	Orders map[string]*engine.Order
	mu     sync.RWMutex
}

// Initialize new memory storage
func NewMemory() *Memory {
	return &Memory{
		Users:  make(map[string]*engine.User),
		Orders: make(map[string]*engine.Order),
	}
}

// Add a user
func (m *Memory) AddUser(u *engine.User) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Users[u.ID] = u
}

// Get a user by ID
func (m *Memory) GetUser(id string) *engine.User {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Users[id]
}

// Add an order
func (m *Memory) AddOrder(o *engine.Order) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Orders[o.ID] = o
}

// Get an order by ID
func (m *Memory) GetOrder(id string) *engine.Order {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Orders[id]
}

// Delete an order
func (m *Memory) DeleteOrder(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Orders, id)
}
