package engine

// OrderHeap is a min-heap or max-heap depending on the comparator
type OrderHeap struct {
	Data []*Order               // stores orders
	Less func(a, b *Order) bool // comparator function (min or max)
}

// NewOrderHeap creates a new heap with a comparator
func NewOrderHeap(less func(a, b *Order) bool) *OrderHeap {
	return &OrderHeap{
		Data: []*Order{},
		Less: less,
	}
}

// Insert adds an order to the heap
func (h *OrderHeap) Insert(o *Order) {
	h.Data = append(h.Data, o) // append at the end
	o.Index = len(h.Data) - 1  // track index
	h.heapifyUp(o.Index)       // restore heap property
}

// PeekRoot returns the top of the heap (best bid/ask)
func (h *OrderHeap) PeekRoot() *Order {
	if len(h.Data) == 0 {
		return nil
	}
	return h.Data[0]
}

// Remove deletes an order from heap
func (h *OrderHeap) Remove(o *Order) {
	index := o.Index
	last := len(h.Data) - 1

	h.swap(index, last)    // swap with last
	h.Data = h.Data[:last] // remove last

	if index < len(h.Data) {
		h.heapifyUp(index) // restore heap
		h.heapifyDown(index)
	}
}

// Update adjusts order position after quantity change
func (h *OrderHeap) Update(o *Order) {
	h.heapifyUp(o.Index)
	h.heapifyDown(o.Index)
}

// heapifyUp restores heap property upwards
func (h *OrderHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.Less(h.Data[i], h.Data[parent]) {
			h.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

// heapifyDown restores heap property downwards
func (h *OrderHeap) heapifyDown(i int) {
	size := len(h.Data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		best := i

		if left < size && h.Less(h.Data[left], h.Data[best]) {
			best = left
		}
		if right < size && h.Less(h.Data[right], h.Data[best]) {
			best = right
		}
		if best == i {
			break
		}
		h.swap(i, best)
		i = best
	}
}

// swap exchanges two elements and updates their index
func (h *OrderHeap) swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
	h.Data[i].Index = i
	h.Data[j].Index = j
}

// BidLess comparator for max-heap (bids)
func BidLess(a, b *Order) bool {
	return a.Price > b.Price
}

// AskLess comparator for min-heap (asks)
func AskLess(a, b *Order) bool {
	return a.Price < b.Price
}
