package leetcode

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 二叉搜索树中第K小的元素
// 思路： 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	var (
		index int
		ans   = -1
	)

	var inOrder func(cur *TreeNode)
	inOrder = func(cur *TreeNode) {
		if cur != nil {
			inOrder(cur.Left)
			index++
			if index == k {
				ans = cur.Val
				return
			}
			inOrder(cur.Right)
		}

	}

	inOrder(root)
	return ans
}

// 进阶：
// 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
// 先中序遍历存储为数组，然后每次插入删除，同步更新数组
