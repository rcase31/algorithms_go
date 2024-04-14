package heapsort

import "testing"

// page176
func TestEx651(t *testing.T) {
	A := heap{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}
	n := len(A)
	maximum, err := MaxHeapExtractMax(&A)
	if err != nil {
		t.Error("should not have error")
	}
	if maximum != 15 {
		t.Error("should be 15")
	}
	nnew := len(A)
	if n == nnew {
		t.Error("should have been reduced")
	}
}

func TestIncreaseKey(t *testing.T) {
	A := heap{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}
	Print(A)
	err := MaxHeapIncreaseKey(A, 2, 50)
	if err != nil {
		t.Error("should not have error")
	}
	Print(A)
}

func TestInsert(t *testing.T) {
	A := heap{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}
	Print(A)
	MaxHeapInsert(&A, 50)
	Print(A)
}
