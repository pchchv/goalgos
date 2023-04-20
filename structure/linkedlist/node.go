package linkedlist

// Node is a structure representing the linkedlist node.
// This node is shared across different implementations.
type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil, nil}
}
