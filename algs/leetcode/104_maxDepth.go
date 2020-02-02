package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightDepth, leftDepth := maxDepth(root.Right), maxDepth(root.Left)

	return max(rightDepth, leftDepth) + 1
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	var tmpNodes []*TreeNode
	for len(nodes) != 0 {
		depth++
		tmpNodes = nil
		for i := 0; i < len(nodes); i++ {
			if nodes[i].Left != nil {
				tmpNodes = append(tmpNodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}
		}

		nodes = tmpNodes
	}

	return depth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
