package heapsort

import (
	"errors"
	"fmt"
	"math"
)

// page 162
func Parent(index int) int {
	return int(math.Floor(float64(index) / 2))
}

func Left(index int) int {
	return index*2 + 1
}

func Right(index int) int {
	return index*2 + 2
}

type Element[T any] struct {
	Key   int
	ID    int
	Value T
}

type Heap[T any] []Element[T]

func NewHeap(ns ...int) Heap[struct{}] {
	h := make([]Element[struct{}], len(ns))
	for i, n := range ns {
		h[i] = Element[struct{}]{Key: n}
	}
	return h
}

func (h Heap[T]) ToList() []int {
	n := len(h)
	out := make([]int, n)
	for i, v := range h {
		out[i] = v.Key
	}
	return out
}

func AssertMaxHeapProperty[T any](A Heap[T], index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex].Key >= A[index].Key
}

func AssertMinHeapProperty[T any](A Heap[T], index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex].Key <= A[index].Key
}

// page 165
// index is a new element inserted into the heap
// we know for sure that tress rooted at left and right are max heaps
func MaxHeapify[T any](A Heap[T], index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	Print(A)

	largest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex].Key > A[index].Key {
		largest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex].Key > A[largest].Key {
		largest = rightIndex
		isUnbalanced = true
	}

	if isUnbalanced {
		// swap of places
		A[largest], A[index] = A[index], A[largest]
		// mind that largest only represents the index of the newly inserted element, since it is no longer the largest
		MaxHeapify(A, largest)
	}
}

func MinHeapify[T any](A Heap[T], index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	fmt.Println("before heapify")
	Print(A)

	smallest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex].Key < A[index].Key {
		smallest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex].Key < A[smallest].Key {
		smallest = rightIndex
		isUnbalanced = true
	}

	if isUnbalanced {
		// swap of places
		A[smallest], A[index] = A[index], A[smallest]
		// mind that largest only represents the index of the newly inserted element, since it is no longer the largest
		MinHeapify(A, smallest)
	}

	fmt.Println("After heapify")
	Print(A)
}

// page 167
// we depart from the upper half of the heap since all bellow are one-sized heaps (leaves)
func BuildMaxHeap[T any](A Heap[T]) {
	n := len(A)
	halfwayIndex := math.Floor(float64(n) / 2)
	for i := halfwayIndex; i >= 0; i-- {
		MaxHeapify(A, int(i))
	}
}

func BuildMinHeap[T any](A Heap[T]) {
	n := len(A)
	halfwayIndex := math.Floor(float64(n) / 2)
	for i := halfwayIndex; i >= 0; i-- {
		MinHeapify(A, int(i))
	}
}

// page 170
func HeapSort[T any](A Heap[T]) {
	BuildMaxHeap(A) // O(n)
	n := len(A)
	for i := n - 1; i > 0; i-- { // O(n)
		A[i], A[0] = A[0], A[i] // A[0] will always be the biggest in reference to the current heapsize
		last := i
		MaxHeapify(A[:last], 0) // O(logn)
	}
}

// not working
func heapSortWithMinHeap[T any](A Heap[T]) {
	BuildMinHeap(A)
	n := len(A)
	for i := 1; i <= n; i++ {
		MinHeapify(A[i:], 0)
	}
}

var ErrHeapUnderflow = errors.New("heap underflow")

func MaxHeapMaximum[T any](A Heap[T]) (*Element[T], error) {
	if A == nil {
		return nil, ErrHeapUnderflow
	}
	return &A[0], nil
}

func MaxHeapExtractMax[T any](A *Heap[T]) (*Element[T], error) {
	maximum, err := MaxHeapMaximum(*A)
	if err != nil {
		return nil, err
	}
	lastIndex := len(*A) - 1
	(*A)[0] = (*A)[lastIndex]
	*A = (*A)[:lastIndex-1]
	MaxHeapify(*A, 0)
	return maximum, nil
}

func MinHeapMinimum[T any](A Heap[T]) (*Element[T], error) {
	if A == nil {
		return nil, ErrHeapUnderflow
	}
	return &A[0], nil
}

func MinHeapExtractMin[T any](A *Heap[T]) (*Element[T], error) {
	minimum, err := MinHeapMinimum(*A)
	if err != nil {
		return nil, err
	}
	lastIndex := len(*A) - 1
	(*A)[0] = (*A)[lastIndex]
	*A = (*A)[:lastIndex]
	MinHeapify(*A, 0)
	return minimum, nil
}

// page 176
// xi is the index of the value to be substituted
// k is the new value

var ErrNewKeySmaller = errors.New("new key is smaller than current key")
var ErrHeapOverflow = errors.New("heap overflow")

func MaxHeapIncreaseKey[T any](A Heap[T], xi, k int) error {
	if A[xi].Key > k {
		return ErrNewKeySmaller
	}
	A[xi].Key = k
	i := xi
	for i > 0 && A[Parent(i)].Key < A[i].Key {
		A[Parent(i)], A[i] = A[i], A[Parent(i)]
		i = Parent(i)
	}
	return nil
}

// xi is the index
func MaxHeapInsert[T any](A *Heap[T], x Element[T]) {
	k := x.Key
	x.Key = math.MinInt
	*A = append(*A, x)
	xi := len(*A) - 1
	// err will never occur
	_ = MaxHeapIncreaseKey(*A, xi, k)
}

func MinHeapDecreaseKey[T any](A Heap[T], xi, k int) error {
	if A[xi].Key > k {
		return ErrNewKeySmaller
	}
	A[xi].Key = k
	i := xi
	for i > 0 && A[Parent(i)].Key > A[i].Key {
		A[Parent(i)], A[i] = A[i], A[Parent(i)]
		i = Parent(i)
	}
	return nil
}

// xi is the index
func MinHeapInsert[T any](A *Heap[T], x Element[T]) {
	k := x.Key
	x.Key = math.MinInt
	*A = append(*A, x)
	xi := len(*A) - 1
	// err will never occur
	_ = MinHeapDecreaseKey(*A, xi, k)
}
