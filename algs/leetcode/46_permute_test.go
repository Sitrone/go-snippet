package leetcode

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)

	fmt.Println(len(ans))
	fmt.Println(ans)
}
