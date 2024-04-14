package heapsort

import (
	"errors"
	"math"
)

var ErrHeapUnderflow = errors.New("heap underflow")

func MaxHeapMaximum(A heap) (int, error) {
	if A == nil {
		return 0, ErrHeapUnderflow
	}
	return A[0], nil
}

func MaxHeapExtractMax(A *heap) (int, error) {
	maximum, err := MaxHeapMaximum(*A)
	if err != nil {
		return 0, err
	}
	lastIndex := len(*A) - 1
	(*A)[0] = (*A)[lastIndex]
	*A = (*A)[:lastIndex-1]
	MaxHeapify(*A, 0)
	return maximum, nil
}

// page 176
// xi is the index of the value to be substituted
// k is the new value

var ErrNewKeySmaller = errors.New("new key is smaller than current key")
var ErrHeapOverflow = errors.New("heap overflow")

func MaxHeapIncreaseKey(A heap, xi, k int) error {
	if A[xi] > k {
		return ErrNewKeySmaller
	}
	A[xi] = k
	i := xi
	for i > 0 && A[Parent(i)] < A[i] {
		A[Parent(i)], A[i] = A[i], A[Parent(i)]
		i = Parent(i)
	}
	return nil
}

// xi is the index
func MaxHeapInsert(A *heap, xv int) {
	k := xv
	xv = math.MinInt
	*A = append(*A, xv)
	xi := len(*A) - 1
	// err will never occur
	_ = MaxHeapIncreaseKey(*A, xi, k)

}

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
