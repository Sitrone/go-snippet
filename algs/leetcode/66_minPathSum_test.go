package leetcode

import (
	"fmt"
	"testing"
)

func TestMinPath(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}
