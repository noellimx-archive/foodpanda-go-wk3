package linkedlist

import "log"

type TLinkedList[T any] struct {
	head *TNode[T]
	size int
}

func (ll *TLinkedList[T]) ResetHead(h_ *TNode[T]) {

	log.Println("(ll *TLinkedList[T]) [ResetHead]")
	ll.head = h_
	curr_ := ll.head
	size_ := 0

	for curr_ != nil {
		size_++
		curr_ = curr_.Next()
	}

	ll.size = size_

	log.Println("(ll *TLinkedList[T]) [ResetHead] head reassigned and size calibrated")

}

func NewLinkedListSingle[T any]() *TLinkedList[T] {
	l := &TLinkedList[T]{
		head: &TNode[T]{},
		size: 0,
	}
	return l
}
