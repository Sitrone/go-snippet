package leetcode

func maxDepthAfterSplit(seq string) []int {
	if len(seq) == 0 {
		return nil
	}

	var idx int
	ans := make([]int, len(seq))
	for _, c := range seq {
		if c == '(' {
			ans[idx] = idx & 0x1
		} else {
			ans[idx] = (idx + 1) & 0x1
		}
		idx++
	}

	return ans
}
