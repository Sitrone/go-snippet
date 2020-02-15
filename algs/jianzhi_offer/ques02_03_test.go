package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 9}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	deleteNode(node2)

	fmt.Println(node1)
}
