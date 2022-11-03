package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type typePrim int

func TestEmptyListD(t *testing.T) {

	ll := NewLinkedListDouble[typePrim]()

	sizeWant := 0

	sizeGot := ll.Size()

	assert.Equal(t, sizeGot, sizeWant, "Empty list should have size 0.")

}
