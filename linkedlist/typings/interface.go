package linkedlist

type list[N LinkedListD[T] | LinkedListS[T], T any] interface {
	ResetHead(*N)
	PrintAll()
}

func CheckListAlike[T any, N LinkedListD[T] | LinkedListS[T]](ll *list[N, T]) {

}
