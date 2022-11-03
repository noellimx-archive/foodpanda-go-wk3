package linkedlist

type NodeD[T any] struct {
	prev *NodeS[T]
	NodeS[T]
}

func (node *NodeD[T]) Val() *T {
	return &node.val
}

func (node *NodeD[T]) Next() *NodeD[T] {
	return node.Next()
}

func (node *NodeD[T]) SetNextNode(next_ *NodeD[T]) {
	node.SetNextNode(next_)
}

func NewNodeD[T any](val_ T) *NodeD[T] {
	n := &NodeD[T]{}

	n.SetVal(val_)
	return n
}
