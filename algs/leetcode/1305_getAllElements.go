package leetcode

// 性质: 二叉搜索树中序遍历为单调递增的有序数组
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := make([]int, 0), make([]int, 0)
	inOrder(root1, &arr1)
	inOrder(root2, &arr2)

	var (
		ans     = make([]int, len(arr1)+len(arr2))
		i, j, k int
	)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			ans[k] = arr1[i]
			i++
		} else {
			ans[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		ans[k] = arr1[i]
		k++
		i++
	}

	for j < len(arr2) {
		ans[k] = arr2[j]
		k++
		j++
	}

	return ans
}

func inOrder(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, ret)
	*ret = append(*ret, root.Val)
	inOrder(root.Right, ret)
}
