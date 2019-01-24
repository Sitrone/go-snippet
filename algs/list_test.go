package algs

import (
	"fmt"
	"testing"
)

func TestListNode_Reverse(t *testing.T) {
	root := &ListNode{Value: 1}
	n1 := &ListNode{Value: 2}
	n2 := &ListNode{Value: 2}
	n3 := &ListNode{Value: 3}
	n4 := &ListNode{Value: 4}

	root.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	rHead := root.Reverse()
	fmt.Println(rHead)
}
