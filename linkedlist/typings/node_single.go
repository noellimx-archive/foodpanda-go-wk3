package linkedlist

type NodeS[T any] struct {
	val  T
	next *NodeS[T]
}

func (node *NodeS[T]) Val() *T {
	return &node.val
}

func (node NodeS[T]) Next() *NodeS[T] {
	return node.next
}

func (node NodeS[T]) SetNextNode(next_ *NodeS[T]) {
	node.next = next_
}

func NewNodeS[T any](val_ T) *NodeS[T] {
	return &NodeS[T]{val: val_}
}
