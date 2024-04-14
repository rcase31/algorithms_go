package heapsort

import (
	"fmt"
	"testing"
)

func TestEX621(t *testing.T) {
	A := NewHeap(27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0)
	MaxHeapify(A, 2)
	Print(A)
}

func TestEX641(t *testing.T) {
	A := NewHeap(5, 13, 2, 25, 7, 17, 20, 8, 4)
	HeapSort(A)
	fmt.Println(A)
}

func TestEX641B(t *testing.T) {
	A := NewHeap(5, 13, 2, 25, 7, 17, 20, 8, 4)
	heapSortWithMinHeap(A)
	fmt.Println(A)
}

// page176
func TestEx651(t *testing.T) {
	A := NewHeap(15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1)
	n := len(A)
	maximum, err := MaxHeapExtractMax(&A)
	if err != nil {
		t.Error("should not have error")
	}
	if maximum.Key != 15 {
		t.Error("should be 15")
	}
	nnew := len(A)
	if n == nnew {
		t.Error("should have been reduced")
	}
}

func TestIncreaseKey(t *testing.T) {
	A := NewHeap(15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1)
	Print(A)
	err := MaxHeapIncreaseKey(A, 2, 50)
	if err != nil {
		t.Error("should not have error")
	}
	Print(A)
}

func TestInsert(t *testing.T) {
	A := NewHeap(15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1)
	Print(A)
	MaxHeapInsert(&A, Element[struct{}]{Key: 50})
	Print(A)
}
