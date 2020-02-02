package leetcode

// 取极值，双指针
func diStringMatch(S string) []int {
	min, max := 0, len(S)

	var ans = make([]int, len(S)+1)
	for i, v := range S {
		if v == 'I' {
			ans[i] = min
			min++
		} else {
			ans[i] = max
			max--
		}
	}

	ans[len(S)] = min
	return ans
}
