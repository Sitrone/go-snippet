package algs

import (
	"testing"
)

func TestPrintListFromEnd2Front(t *testing.T) {
	head := NewList(1, 2, 3, 4, 5)
	PrintListFromEnd2Front(head)

	PrintListFromEnd2Front1(head)
}
