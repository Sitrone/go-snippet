package leetcode

import "math"

// 错误解法 应该是左子树的整颗子树节点值都小于当前节点，而不只是当前节点的左子树节点
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val > root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val < root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

// 中序遍历
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		inorder func(cur *TreeNode)
		isValid = true
		curVal  = math.MinInt64
	)
	inorder = func(cur *TreeNode) {
		if cur == nil || !isValid {
			return
		}
		inorder(cur.Left)
		if cur.Val <= curVal {
			isValid = false
		} else {
			curVal = cur.Val
		}
		inorder(cur.Right)
	}

	inorder(root)
	return isValid
}
