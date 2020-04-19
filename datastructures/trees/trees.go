package trees

// MultiNode multi node
type MultiNode struct {
	Parent   *MultiNode
	Data     interface{}
	Children []*MultiNode
}

// BinaryNode binary node
type BinaryNode struct {
	Parent *BinaryNode
	Data   interface{}
	Left   *BinaryNode
	Right  *BinaryNode
}
