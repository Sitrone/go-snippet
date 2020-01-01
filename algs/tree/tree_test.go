package tree

import (
	"fmt"
	"testing"
)

func TestPreTravers(t *testing.T) {
	root := &Node{val: 1}
	root.left = &Node{val: 2}
	root.right = &Node{val: 3}
	left := root.left
	left.left = &Node{val: 4}
	right := root.right
	right.left = &Node{val: 5}
	right.right = &Node{val: 6}

	rright := right.right
	rright.left = &Node{val: 7}

	//preTraverse(root)
	//fmt.Println()
	//preTraverseIter(root)

	inTraverse(root)
	fmt.Println()
	inTraverseIter(root)

	m := make(map[string]int)
	fmt.Println(len(m))

	var s []int
	fmt.Println(len(s))
}
