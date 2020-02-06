package leetcode

func plusOne(digits []int) []int {
	var (
		i, sum, canary, first int
		ans                   = make([]int, 0, len(digits)+1)
	)

	first = 1
	for i = len(digits) - 1; i >= 0; i, first = i-1, 0 {
		sum = canary + digits[i] + first
		sum, canary = sum%10, sum/10
		ans = append(ans, sum)
	}

	// 避免使用ans = append([]int{canary}, ans...) 性能差
	if canary > 0 {
		ans = append(ans, canary)
	}

	for j := 0; j < len(ans)/2; j++ {
		ans[j], ans[len(ans)-j-1] = ans[len(ans)-j-1], ans[j]
	}

	return ans
}
