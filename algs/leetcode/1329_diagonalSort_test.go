package leetcode

import (
	"fmt"
	"testing"
)

func TestDiagonalSort(t *testing.T) {
	matrix := [][]int{
		{3, 3, 1, 1},
		{2, 2, 1, 2},
		{1, 1, 1, 2},
	}

	fmt.Println(diagonalSort(matrix))
}
