package leetcode

// 广度优先，层序遍历，计算每一层的值
func deepestLeavesSum(root *TreeNode) int {
	var (
		queue []*TreeNode
		total int
	)

	queue = append(queue, root)
	for len(queue) != 0 {
		total = 0
		var tmpQueue []*TreeNode
		for _, node := range queue {
			total += node.Val
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}

		queue = tmpQueue
	}

	return total
}
