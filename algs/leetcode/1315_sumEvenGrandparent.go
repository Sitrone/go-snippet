package leetcode

// 可以计算当前节点如果为偶数的所有孙子节点
func sumEvenGrandparent(root *TreeNode) int {
	var (
		sum   int
		queue []*TreeNode
	)
	queue = append(queue, root)

	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]

		if head.Val&0x01 == 0 {
			sum += sumEven(head)
		}
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}

	return sum
}

// 递归解法
func sumEvenGrandparent1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	if root.Val&0x01 == 0 {
		sum += sumEven(root)
	}

	if root.Left != nil {
		sum += sumEvenGrandparent1(root.Left)
	}

	if root.Right != nil {
		sum += sumEvenGrandparent1(root.Right)
	}

	return sum
}

func sumEven(root *TreeNode) int {
	var sum int
	if root.Left != nil {
		if root.Left.Left != nil {
			sum += root.Left.Left.Val
		}
		if root.Left.Right != nil {
			sum += root.Left.Right.Val
		}
	}

	if root.Right != nil {
		if root.Right.Left != nil {
			sum += root.Right.Left.Val
		}
		if root.Right.Right != nil {
			sum += root.Right.Right.Val
		}
	}

	return sum
}
