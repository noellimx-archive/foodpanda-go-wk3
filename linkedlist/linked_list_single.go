package linkedlist

type TLinkedList[T any] struct {
	head *TNode[T]
	size int
}

func (ll *TLinkedList[T]) SetHead(h_ *TNode[T]) {
	ll.head = h_

	curr_ := ll.head

	size_ := 0

	for curr_ != nil {
		size_++
		curr_ = curr_.Next()
	}

	ll.size = size_
}

func NewLinkedListSingle[T any]() *TLinkedList[T] {
	l := &TLinkedList[T]{
		head: &TNode[T]{},
		size: 0,
	}
	return l
}
