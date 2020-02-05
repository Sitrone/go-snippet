package leetcode

// dp dp[0] = 0, dp[1] = a[1]-a[0](if a[1]>a[0])
// dp[i] = dp[i-1] ()
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var (
		dp       = make([]int, len(prices)+1)
		ans, min int
	)

	dp[0] = 0
	min = prices[0]
	for i := 1; i < len(prices); i++ {
		if i > 1 {
			if prices[i-1] < min {
				min = prices[i-1]
			}
		}
		if prices[i] >= prices[i-1] {
			dp[i] = prices[i] - min
		} else {
			dp[i] = dp[i-1]
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
