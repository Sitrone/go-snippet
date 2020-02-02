package leetcode

// [1,null,1,1,1,1,1,1]  TODO HARD
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var rootLen int
	nodeUnivaluePath(root, &rootLen)
	return rootLen
}

func nodeUnivaluePath(root *TreeNode, count *int) int {
	if root == nil {
		return 0
	}

	left, right := nodeUnivaluePath(root.Left, count), nodeUnivaluePath(root.Right, count)

	if root.Left != nil && root.Left.Val == root.Val {
		left += 1
	} else {
		left = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		right += 1
	} else {
		right = 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	*count = max(*count, left+right)
	return max(left, right)
}
