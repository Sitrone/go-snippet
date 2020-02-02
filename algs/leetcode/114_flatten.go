package leetcode

// 层序遍历 TODO
func flatten(root *TreeNode) {
	var ans = make([]int, 0)

	var preOrder = func(root *TreeNode) {}

	preOrder = func(head *TreeNode) {
		ans = append(ans, head.Val)
		preOrder(head.Left)
		preOrder(head.Right)
	}

	preOrder(root)

}
