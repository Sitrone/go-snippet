package jianzhi_offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right, left := mirrorTree(root.Right), mirrorTree(root.Left)
	root.Left, root.Right = right, left
	return root
}
