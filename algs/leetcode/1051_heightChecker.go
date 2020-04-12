package leetcode

import "sort"

func heightChecker(heights []int) int {
	if len(heights) < 2 {
		return 0
	}

	tmp := make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)

	var ret int
	for i := 0; i < len(heights); i++ {
		if tmp[i] != heights[i] {
			ret++
		}
	}

	return ret
}
