package heapsort

// page 173
type PriorityQueue[T any] interface {
	Insert(T, int, int)
	Minimum() *Element[T]
	ExtractMinimum() *Element[T]
	DecreaseKey(int, int)
	IsEmpty() bool
}

type priorityQueue[T any] struct {
	A Heap[T]
}

func NewPQ[T any](e ...Element[T]) PriorityQueue[T] {
	return &priorityQueue[T]{A: e}
}

func (q *priorityQueue[T]) IsEmpty() bool {
	return len(q.A) == 0
}

func (q *priorityQueue[T]) Insert(v T, k int, id int) {
	x := Element[T]{
		ID:    id,
		Key:   k,
		Value: v,
	}
	MinHeapInsert(&q.A, x)
}

func (q *priorityQueue[T]) Minimum() *Element[T] {
	minimum, err := MinHeapMinimum(q.A)
	if err != nil {
		return nil
	}
	return minimum
}

func (q *priorityQueue[T]) ExtractMinimum() *Element[T] {
	maximum, err := MinHeapExtractMin(&q.A)
	if err != nil {
		return nil
	}
	return maximum
}

func (q *priorityQueue[T]) DecreaseKey(xID int, k int) {
	var xi int
	for i, e := range q.A {
		if e.ID == xID {
			xi = i
		}
	}
	//TODO: throw error for when not found
	MinHeapDecreaseKey(q.A, xi, k)
}
