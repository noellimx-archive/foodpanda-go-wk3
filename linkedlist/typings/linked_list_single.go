package linkedlist

import (
	"fmt"
	"log"
	"strings"
)

type LinkedListS[ValT any] LinkedList_[ValT]

type LinkedList_[ValT any] struct {
	head *NodeS[ValT]
	tail *NodeS[ValT]
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

func (ll *LinkedListS[T]) sanityCheck() {
	if ll.Size() == 0 || ll.head == nil || ll.tail == nil {
		if !(ll.Size() == 0 && ll.head == nil && ll.tail == nil) {
			log.Fatalln("(ll *LinkedListS[T]) [sanityCheck] failed. Initial state corrupted.")
		}
	}
}

// O(n) operation
func (ll *LinkedListS[T]) RemoveTail() {
	ll.sanityCheck()

	this := ll.head

	switch ll.Size() {
	case 0:
		return
	case 1:
		ll.head = nil
		ll.tail = nil
		ll.size--

		return
	}

	for this != nil {
		next := this.Next()
		if next == nil {

			ll.tail = this
			ll.tail.SetNextNode(nil)
			ll.size--
			return
		}

		this = next
	}

}

func (ll *LinkedListS[T]) InsertNode(next_ *NodeS[T]) {

	if ll.size == 0 {
		ll.ResetHead(next_)
		return
	}
	ll.tail.SetNextNode(next_)
	ll.tail = next_

	ll.size++
}

func (ll *LinkedListS[T]) Size() int {

	log.Println("(ll *TLinkedList[T]) [Size]")

	return ll.size

}

func (ll *LinkedListS[T]) Head() *NodeS[T] {

	log.Println("(ll *TLinkedList[T]) [Head]")

	return ll.head

}

func (ll *LinkedListS[T]) Tail() *NodeS[T] {

	log.Println("(ll *TLinkedList[T]) [Tail]")

	return ll.tail

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
