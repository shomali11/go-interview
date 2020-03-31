package priorityqueues

import "container/heap"

/*
	https://golang.org/pkg/container/heap/
*/

// A heapArray implements heap.Interface and holds Items.
type heapArray struct {
	max   bool
	array []*PQElement
}

// Len length of the queue
func (h heapArray) Len() int { return len(h.array) }

// Less compares priority of two elements in the queue
func (h heapArray) Less(i, j int) bool {
	if h.max {
		return h.array[i].Priority > h.array[j].Priority
	}
	return h.array[i].Priority < h.array[j].Priority
}

// Swap swaps two elements in the queue
func (h heapArray) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
	h.array[i].index = i
	h.array[j].index = j
}

// Push pushes to the queue
func (h *heapArray) Push(x interface{}) {
	n := len(h.array)
	item := x.(*PQElement)
	item.index = n
	h.array = append(h.array, item)
}

// Pop pops from the queue
func (h *heapArray) Pop() interface{} {
	old := h.array
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.array = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (h *heapArray) update(element *PQElement, value string, priority int) {
	element.Value = value
	element.Priority = priority
	heap.Fix(h, element.index)
}
