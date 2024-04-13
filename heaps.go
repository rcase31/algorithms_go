package heapsort

import (
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

type heap []int

func AssertMaxHeapProperty(A heap, index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex] >= A[index]
}

func AssertMinHeapProperty(A heap, index int) bool {
	parentIndex := Parent(index)
	return A[parentIndex] <= A[index]
}

// page 165
// index is a new element inserted into the heap
// we know for sure that tress rooted at left and right are max heaps
func MaxHeapify(A heap, index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	Print(A)

	largest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex] > A[index] {
		largest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex] > A[largest] {
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
