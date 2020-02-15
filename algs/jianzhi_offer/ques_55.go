package jianzhi_offer

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	return 1 + maxTwo(maxDepth(root.Left), maxDepth(root.Right))
}
