package heap_test

import (
	. "github.com/htanata/heap"

	"testing"
)

func TestHeap(t *testing.T) {
	heap := Heap{}

	for _, n := range []int{1, 4, 5, 2, 3} {
		heap.Insert(n)
	}

	for _, n := range []int{1, 2, 3, 4, 5} {
		result := heap.ExtractMin()

		if result != n {
			t.Error("Expected", n, "got", result)
		}
	}
}
