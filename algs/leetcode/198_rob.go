package leetcode

// 基础常规dp问题
// dp[0] = nums[0], dp[1] = max(dp[0], nums[1])
// dp[i] = max(dp[i-1], dp[i-1]-nums[i-1]+nums[i]) if rob last else dp[i] = dp[i-1]+ nums[i]
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curMax, preMax := nums[0], 0
	for i := 1; i < len(nums); i++ {
		temp := curMax
		if preMax+nums[i] > curMax {
			curMax = preMax + nums[i]
		}
		preMax = temp
	}

	return curMax
}
