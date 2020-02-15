package jianzhi_offer

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) String() string {
	return fmt.Sprintf("TreeNode{value=%d, Left=%v, Right-%v}", root.Val, root.Left, root.Right)
}
