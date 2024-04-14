package heapsort

import (
	"fmt"
	"testing"
)

func TestEX621(t *testing.T) {
	A := heap{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	MaxHeapify(A, 2)
	Print(A)
}

func TestEX641(t *testing.T) {
	A := []int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	HeapSort(A)
	fmt.Println(A)
}

func TestEX641B(t *testing.T) {
	A := []int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	heapSortWithMinHeap(A)
	fmt.Println(A)
}
