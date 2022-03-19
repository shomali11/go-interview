package priorityqueues

/*
	https://golang.org/pkg/container/heap/
*/

// pqElement is an value in a priority queue.
type pqElement[T any] struct {
	value T

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A heapArray implements heap.Interface and holds Items.
type heapArray[T any] struct {
	compare func(i, j T) bool
	array   []*pqElement[T]
}

// Len length of the queue
func (h heapArray[T]) Len() int { return len(h.array) }

// Less compares priority of two elements in the queue
func (h heapArray[T]) Less(i, j int) bool {
	return h.compare(h.array[i].value, h.array[j].value)
}

// Swap swaps two elements in the queue
func (h heapArray[T]) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
	h.array[i].index = i
	h.array[j].index = j
}

// Push pushes to the queue
func (h *heapArray[T]) Push(x any) {
	n := len(h.array)
	item := x.(*pqElement[T])
	item.index = n
	h.array = append(h.array, item)
}

// Pop pops from the queue
func (h *heapArray[T]) Pop() any {
	old := h.array
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	h.array = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (h *heapArray) update(element *pqElement, value string, priority int) {
// 	element.Value = value
// 	element.Priority = priority
// 	heap.Fix(h, element.index)
// }
