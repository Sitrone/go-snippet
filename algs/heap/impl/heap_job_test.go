package impl

import (
	"fmt"
	"testing"
)

func TestJobHeap_Pop(t *testing.T) {
	hp := &JobHeap{}

	fmt.Println(hp.Len() == 0)
}
