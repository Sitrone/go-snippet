package leetcode

import (
	"fmt"
	"testing"
)

func TestFindDuplidate(t *testing.T) {
	var nums = []int{1, 3, 4, 2, 1}
	fmt.Println(findDuplicate1(nums))
}
