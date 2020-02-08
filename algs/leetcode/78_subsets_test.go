package leetcode

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}

	ans := subsets(nums)
	fmt.Println(len(ans))
	fmt.Println(ans)
}
