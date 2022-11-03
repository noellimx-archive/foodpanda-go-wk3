package linkedlist

type NodeS[T any] struct {
	val  T
	next *NodeS[T]
}

func (node *NodeS[ValT]) SetVal(val_ ValT) {
	node.val = val_
}

func (node *NodeS[ValT]) Val() *ValT {
	return &node.val
}

func (node *NodeS[ValT]) Next() *NodeS[ValT] {
	return node.next
}

func (node *NodeS[ValT]) SetNextNode(next_ *NodeS[ValT]) {
	node.next = next_
}

func NewNodeS[ValT any](val_ ValT) *NodeS[ValT] {
	return &NodeS[ValT]{val: val_}
}
