package algs

// 重建二叉树， 不含重复数字
func ConstructTree(pre, in []int) *TreeNode {
	return ReConstruct(pre, 0, len(pre)-1, in, 0, len(in)-1)
}

func ReConstruct(pre []int, pStart, pEnd int, in []int, iStart, iEnd int) *TreeNode {
	if pStart > pEnd || iStart > iEnd {
		return nil
	}

	root := TreeNode{Val: pre[pStart]}
	// 需要使用遍历的手段找到in树的root节点位置
	for i := iStart; i <= iEnd; i++ {
		if pre[pStart] == in[i] {
			root.left = ReConstruct(pre, pStart+1, pStart+i-iStart, in, iStart, i-1)
			root.right = ReConstruct(pre, pStart+i-iStart+1, pEnd, in, i+1, iEnd)
		}
	}
	return &root
}
