package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flowe", "flow", ""}
	//strs := []string{"dog", "racecar", "car"}
	ans := longestCommonPrefix1(strs)
	fmt.Println(ans)
}
