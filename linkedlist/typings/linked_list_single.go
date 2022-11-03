package linkedlist

import (
	"fmt"
	"log"
	"strings"
)

type LinkedListS[T any] LinkedList_[T]

type LinkedList_[T any] struct {
	head *NodeS[T]
	tail *NodeS[T]
	size int
}

func (ll *LinkedListS[T]) ResetHead(h_ *NodeS[T]) {

	log.Println("(ll *TLinkedList[T]) [ResetHead]")
	ll.head = h_
	ll.tail = h_
	curr_ := ll.head
	size_ := 0

	for curr_ != nil {
		size_++
		curr_ = curr_.Next()
	}

	ll.size = size_

	log.Println("(ll *TLinkedList[T]) [ResetHead] head reassigned and size calibrated")

}

func (ll *LinkedListS[T]) PrintAll() {

	log.Println("(ll *TLinkedList[T]) [PrintAll]")
	this := ll.head

	var string_ strings.Builder
	for this != nil {
		stringified_val := fmt.Sprintf("%#v", this.Val())

		string_.Write([]byte(stringified_val))
		this = ll.head
	}

	log.Println(this.Val())
}

func NewLinkedListSingle[T any]() *LinkedListS[T] {
	l := &LinkedListS[T]{
		head: &NodeS[T]{},
		size: 0,
	}

	return l
}
