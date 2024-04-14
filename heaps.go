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

type element[T any] struct {
	key   int
	value T
}

type heap[T any] []element[T]

func AssertMaxHeapProperty[T any](A heap[T], index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex].key >= A[index].key
}

func AssertMinHeapProperty[T any](A heap[T], index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex].key <= A[index].key
}

// page 165
// index is a new element inserted into the heap
// we know for sure that tress rooted at left and right are max heaps
func MaxHeapify[T any](A heap[T], index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	Print(A)

	largest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex].key > A[index].key {
		largest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex].key > A[largest].key {
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

func MinHeapify[T any](A heap[T], index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	fmt.Println("before heapify")
	Print(A)

	smallest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex].key < A[index].key {
		smallest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex].key < A[smallest].key {
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
func BuildMaxHeap(A []int) {
	n := len(A)
	halfwayIndex := math.Floor(float64(n) / 2)
	for i := halfwayIndex; i >= 0; i-- {
		MaxHeapify(A, int(i))
	}
}

func BuildMinHeap(A []int) {
	n := len(A)
	halfwayIndex := math.Floor(float64(n) / 2)
	for i := halfwayIndex; i >= 0; i-- {
		MinHeapify(A, int(i))
	}
}

// page 170
func HeapSort(A []int) {
	BuildMaxHeap(A) // O(n)
	n := len(A)
	for i := n - 1; i > 0; i-- { // O(n)
		A[i], A[0] = A[0], A[i] // A[0] will always be the biggest in reference to the current heapsize
		last := i
		MaxHeapify(A[:last], 0) // O(logn)
	}
}

// not working
func heapSortWithMinHeap(A []int) {
	BuildMinHeap(A)
	n := len(A)
	for i := 1; i <= n; i++ {
		MinHeapify(A[i:], 0)
	}
}

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
