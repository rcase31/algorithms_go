package heapsort

import (
	"fmt"
	"testing"
)

func TestEX621(t *testing.T) {
	A := heap{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	MaxHeapify(A, 2)
	Print(A)
	fmt.Println("aa")
}
