package tree

import (
	"container/list"
	"fmt"
	"zitrone.github.com/go-snippet/algs"
)

type Node struct {
	val         int
	left, right *Node
}

func (n *Node) setLeft(left *Node) {
	n.left = left
}

func (n *Node) setRight(right *Node) {
	n.right = right
}

func preTraverse(root *Node) {
	if root != nil {
		fmt.Print(root.val, " ")
		preTraverse(root.left)
		preTraverse(root.right)
	}
}

func preTraverseIter(root *Node) {
	if root == nil {
		return
	}

	var stack algs.Stack
	stack.Push(root)
	for !stack.IsEmpty() {
		if cur, err := stack.Pop(); err != nil {
			panic(err)
		} else {
			if n := cur.(*Node); n != nil {
				fmt.Print(n.val, " ")
				stack.Push(n.right)
				stack.Push(n.left)
			}
		}
	}
}

func inTraverse(root *Node) {
	if root != nil {
		inTraverse(root.left)
		fmt.Print(root.val, " ")
		inTraverse(root.right)
	}
}

func inTraverseIter(root *Node) {
	t := root
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.left
		}

		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*Node)
			res = append(res, t.val) //visit
			t = t.right
			stack.Remove(v)
		}
	}

	fmt.Println(res)
}

func postTraverse(root *Node) {
	if root != nil {
		postTraverse(root.left)
		postTraverse(root.right)
		fmt.Println(root.val)
	}
}
