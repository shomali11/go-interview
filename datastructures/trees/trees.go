package trees

// MultiNode multi node
type MultiNode[T any] struct {
	Parent   *MultiNode[T]
	Data     T
	Children []*MultiNode[T]
}

// BinaryNode binary node
type BinaryNode[T any] struct {
	Parent *BinaryNode[T]
	Data   T
	Left   *BinaryNode[T]
	Right  *BinaryNode[T]
}
