package jianzhi_offer

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		tmp := make([]int, 0, len(curLevel)*3/2)
		nextLevel := make([]*TreeNode, 0)
		for _, t := range curLevel {
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				nextLevel = append(nextLevel, t.Left)
			}
			if t.Right != nil {
				nextLevel = append(nextLevel, t.Right)
			}
		}

		ans = append(ans, tmp)
		curLevel = nextLevel
	}

	return ans
}
