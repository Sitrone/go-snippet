package leetcode

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 6}
	fmt.Println(findKthLargest(nums, 2))
}
