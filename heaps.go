package heapsort

import (
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

func MinHeapify(A heap, index int) {
	leftIndex := Left(index)
	rightIndex := Right(index)
	heapSize := len(A)

	fmt.Println("before heapify")
	Print(A)

	smallest := index
	isUnbalanced := false

	if leftIndex < heapSize && A[leftIndex] < A[index] {
		smallest = leftIndex
		isUnbalanced = true
	}

	if rightIndex < heapSize && A[rightIndex] < A[smallest] {
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
