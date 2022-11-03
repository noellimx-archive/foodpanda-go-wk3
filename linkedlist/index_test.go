package linkedlist

import (
	"fmt"
	"testing"

	linkedlist "github.com/noellimx/foodpanda-go-wk3/linkedlist/typings"
	"github.com/stretchr/testify/assert"
)

type typePrim int

func newEmptyListSingleInts() *linkedlist.LinkedListS[typePrim] {
	return linkedlist.NewLinkedListSingle[typePrim]()
}
func TestNewEmptyListS(t *testing.T) {

	listInts := newEmptyListSingleInts()

	sizeWant := 0
	sizeGot := listInts.Size()
	assert.Equal(t, sizeGot, sizeWant, "Empty list should have size 0.")
}

func TestAddOneToListS(t *testing.T) {

	ll := newEmptyListSingleInts()
	node0 := linkedlist.NewNodeS[typePrim](1)

	ll.ResetHead(node0)

	sizeWant := 1
	sizeGot := ll.Size()
	assert.Equal(t, sizeGot, sizeWant, "list should have size 1.")

	tailWant := node0
	tailGot := ll.Tail()
	assert.Equal(t, tailGot, tailWant, "given size 1 tailGot should be tailWant.")

}

func TestAddTwoToListS(t *testing.T) {

	ll := newEmptyListSingleInts()
	node0 := linkedlist.NewNodeS[typePrim](1)
	node1 := linkedlist.NewNodeS[typePrim](2)

	ll.ResetHead(node0)

	ll.InsertNode(node1)

	sizeWant := 2
	sizeGot := ll.Size()
	assert.Equal(t, sizeWant, sizeGot, fmt.Sprintf("given size 2 should be 2. %d %d", sizeWant, sizeGot))

	tailWant := node1
	tailGot := ll.Tail()

	assert.Equal(t, tailGot, tailWant, "given size 2 tailGot should be tailWant.")

}
