package heapsort

// page 173
type Element[T any] struct {
	Key   int
	Value T
}
type PriorityQueue[T any] interface {
	Insert(Element[T], int)
	Maximum() Element[T]
	ExtractMaximum() Element[T]
	IncreaseKey(Element[T], int)
}

type priorityQueue[T any] struct {
	A []Element[T]
}

func (q *PriorityQueue[T]) Insert(x Element[T], k int) {
	MaxHeapInsert(A)
}
