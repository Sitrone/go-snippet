package leetcode

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	ans := subsetsWithDup(nums)

	fmt.Println(len(ans))
	fmt.Println(ans)
}
