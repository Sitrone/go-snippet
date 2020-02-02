package leetcode

import (
	"fmt"
	"testing"
)

func TestAllPathSource(t *testing.T) {
	var graph = [][]int{{1, 2}, {3}, {3}, {}}
	ret := allPathsSourceTarget(graph)
	fmt.Println(ret)
}
