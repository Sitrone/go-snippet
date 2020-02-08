package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	var candidate = []int{10, 1, 2, 7, 6, 1, 5}

	fmt.Println(combinationSum2(candidate, 8))

	sort.Ints(candidate)
	fmt.Println(candidate)
}
