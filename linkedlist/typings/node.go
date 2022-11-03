package linkedlist

type TNode[T any] struct {
	val  T
	next *TNode[T]
}

func (node *TNode[T]) Val() *T {
	return &node.val
}

func (node TNode[T]) Next() *TNode[T] {
	return node.next
}

func (node TNode[T]) SetNextNode(next_ *TNode[T]) {
	node.next = next_
}

func NewNode[T any](val_ T) *TNode[T] {
	return &TNode[T]{val: val_}
}
