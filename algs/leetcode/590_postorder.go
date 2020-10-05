package leetcode

import "zitrone.github.com/go-snippet/algs"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 给定一个 N 叉树，返回其节点值的后序遍历。
// 递归
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret = make([]int, 0, 10)
	for i := 0; i < len(root.Children); i++ {
		ret = append(ret, postorder(root.Children[i])...)
	}

	ret = append(ret, root.Val)
	return ret
}

// 非递归 stack
func postorder1(root *Node) []int {
	if root == nil {
		return nil
	}
	var (
		stack = &algs.Stack{}
		ret   = make([]int, 0, 10)
	)
	stack.Push(root.Val)
	for !stack.IsEmpty() {
		rootI, _ := stack.Pop()
		root := rootI.(*Node)
		if root != nil {
			ret = append(ret, root.Val)
			for i := 0; i < len(root.Children); i++ {
				stack.Push(root.Children[i])
			}
		}
	}

	for i, j := 0, len(ret)-1; i < len(ret)/2; i++ {
		ret[i], ret[j] = ret[j], ret[i]
		j--
	}
	return ret
}
