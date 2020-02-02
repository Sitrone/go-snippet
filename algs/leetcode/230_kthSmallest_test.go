package leetcode

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	fmt.Println(kthSmallest(root, 1))
}
