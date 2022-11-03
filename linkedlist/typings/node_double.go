package linkedlist

type NodeD[T any] struct {
	prev *NodeD[T]
	val  T
	next *NodeD[T]
}

func (node *NodeD[T]) Val() *T {
	return &node.val
}

func (node *NodeD[T]) Next() *NodeD[T] {
	return node.next
}

func (node *NodeD[T]) SetNextNode(next_ *NodeD[T]) {
	node.SetNextNode(next_)
}

func NewNodeD[T any](val_ T) *NodeD[T] {
	n := &NodeD[T]{val: val_}

	return n
}
