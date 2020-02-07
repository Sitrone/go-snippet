package jianzhi_offer

import "fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func (root *TreeNode) String() string {
	return fmt.Sprintf("TreeNode{value=%d, left=%v, right-%v}", root.Val, root.left, root.right)
}
