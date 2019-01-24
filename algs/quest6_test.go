package algs

import (
	"fmt"
	"testing"
)

func TestConstructTree(t *testing.T) {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}

	root := ConstructTree(pre, in)

	var s []string
	fmt.Println("empty slice is:", s == nil)
	s = append(s, "1")
	fmt.Println(s)

	fmt.Println(root)
}
