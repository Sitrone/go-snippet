package leetcode

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	var arr = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(arr)
	fmt.Println(arr)
}
