package leetcode

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	fmt.Println(rotateRight1(head, 3))
}
