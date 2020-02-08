package leetcode

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 3}
	ans := permuteUnique(nums)

	fmt.Println(len(ans))
	fmt.Println(ans)
}
