package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var list1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	var list2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	var list3 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}

	head := mergeKLists1([]*ListNode{list1, list2, list3})
	fmt.Println(head)
}
