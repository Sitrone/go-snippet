package leetcode

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	var nums = []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
	fmt.Println(rob1(nums))
	fmt.Println(rob2(nums))
}
