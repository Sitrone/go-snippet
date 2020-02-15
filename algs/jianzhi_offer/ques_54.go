package jianzhi_offer

// 中序遍历
func kthLargest(root *TreeNode, k int) int {
	var nums = make([]int, 0, 10)
	var inorder func(cur *TreeNode)
	inorder = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		inorder(cur.Left)
		nums = append(nums, cur.Val)
		inorder(cur.Right)
	}

	inorder(root)
	return nums[len(nums)-k]
}
