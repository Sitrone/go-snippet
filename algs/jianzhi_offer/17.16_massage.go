package jianzhi_offer

// dp[i][j], j = 0, 第i个接单，j=1第i个不接单
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var dp = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i][0] = maxTwo(dp[i-1][0], dp[i-1][1]) // 不接单，前一个可接可不接
		dp[i][1] = dp[i-1][0] + nums[i]           // 接了，前面一个肯定不能接
	}

	return maxTwo(dp[len(nums)-1][0], dp[len(nums)-1][1])
}
