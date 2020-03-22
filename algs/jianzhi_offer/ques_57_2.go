package jianzhi_offer

func findContinuousSequence(target int) [][]int {
	reminder := func(n int) int {
		return (1 + n - 1) * (n - 1) / 2
	}

	ans := make([][]int, 0, 10)
	for i := target / 2; i > 1; i-- {
		total := target - reminder(i)
		if total <= 0 || total%i != 0 {
			continue
		}

		start := total / i
		cur := make([]int, i)
		for j := 0; j < i; j++ {
			cur[j] = start + j
		}

		ans = append(ans, cur)
	}

	return ans
}
