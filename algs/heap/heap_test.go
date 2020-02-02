package heap

import (
	"fmt"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap(10)
	maxHeap.Insert(5)
	maxHeap.Insert(3)
	maxHeap.Insert(7)
	maxHeap.Insert(2)
	maxHeap.Insert(12)
	maxHeap.Insert(6)

	fmt.Println(maxHeap.eles)

	var (
		err error
		top int
	)
	for err == nil {
		top, err = maxHeap.DelTop()
		fmt.Println(top)
	}
}
