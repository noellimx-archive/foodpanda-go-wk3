package linkedlist

import (
	"fmt"
	"log"
	"strings"
)

type LinkedListD[T any] LinkedListS[T]

func (ll *LinkedListD[T]) ResetHead(h_ *NodeS[T]) {

	log.Println("(ll *LinkedListS[T]) [ResetHead]")
	ll.head = h_
	ll.tail = h_
	curr_ := ll.head
	size_ := 0

	for curr_ != nil {
		size_++
		curr_ = curr_.Next()
	}

	ll.size = size_

	log.Println("(ll *LinkedListS[T]) [ResetHead] head reassigned and size calibrated")

}

func (ll *LinkedListD[T]) PrintAll() {

	log.Println("(ll *LinkedListS[T]) [PrintAll]")
	this := ll.head

	var string_ strings.Builder
	for this != nil {
		stringified_val := fmt.Sprintf("%#v", this.Val())

		string_.Write([]byte(stringified_val))
		this = ll.head
	}

	log.Println(this.Val())
}

func NewLinkedListDouble[T any]() *LinkedListD[T] {
	l := &LinkedListD[T]{
		head: &NodeS[T]{},
		size: 0,
	}
	return l
}
