package leetcode

// dfs
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var dfs func(root *TreeNode, sum int) bool
	dfs = func(root *TreeNode, sum int) bool {
		if root.Left == nil && root.Right == nil {
			if sum-root.Val == 0 {
				return true
			} else {
				return false
			}
		}
		var res1, res2 bool
		if root.Left != nil {
			res1 = dfs(root.Left, sum-root.Val)
		}
		if root.Right != nil {
			res2 = dfs(root.Right, sum-root.Val)
		}
		return res1 || res2
	}

	return dfs(root, sum)
}

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
