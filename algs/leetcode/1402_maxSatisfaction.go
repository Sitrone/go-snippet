package leetcode

import "sort"

// 动态规划, dp[i][j]表示前i个物体中，选中j个菜，可以达到的最大值 FIXME 该题解法有问题
// dp[0][1] = arr[0]
// dp[i][j] = max(dp[i-1][j-1], dp[i-1][j] + arr[i]*j)
func maxSatisfaction(satisfaction []int) int {
	var (
		dp      = make([][]int, len(satisfaction))
		maxFunc func(x, y int) int
	)

	sort.Ints(satisfaction)
	maxFunc = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < len(satisfaction); i++ {
		dp[i] = make([]int, len(satisfaction))
	}

	dp[0][0] = 0
	dp[0][1] = satisfaction[0]
	for i := 1; i < len(satisfaction); i++ {
		for j := 0; j < i; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				if i == j {
					dp[i][j] = dp[i-1][j-1] + satisfaction[i-1]*j
				} else {
					dp[i][j] = maxFunc(dp[i-1][j], dp[i-1][j-1]+satisfaction[i-1]*j)
				}
			}
		}
	}

	var ans = 0
	for i := 0; i <= len(satisfaction); i++ {
		ans = maxFunc(ans, dp[len(satisfaction)][i])
	}
	return ans
}
