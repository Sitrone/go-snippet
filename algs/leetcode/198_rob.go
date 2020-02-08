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

// dp
// 使用二维数据记录状态更容易理解dp[i][j], i：打劫到第i家，j=0，第i家不打劫，j=1，第i家打劫
func rob1(nums []int) int {
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
		dp[i][0] = maxTwo(dp[i-1][0], dp[i-1][1]) // 不偷，前一个可偷可不偷
		dp[i][1] = dp[i-1][0] + nums[i]           // 偷了，前面一个肯定不能偷
	}

	return maxTwo(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

// 对于dp 二维数组的优化
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var (
		preNoRob, preRob, max int
		curNoRob, curRob      int
	)

	preNoRob = 0
	preRob = nums[0]
	max = maxTwo(preNoRob, preRob)
	for i := 1; i < len(nums); i++ {
		curNoRob = maxTwo(preNoRob, preRob) // 不偷，前一个可偷可不偷
		curRob = preNoRob + nums[i]         // 偷了，前面一个肯定不能偷

		preNoRob = curNoRob
		preRob = curRob
		max = maxTwo(maxTwo(preNoRob, preRob), max)
	}

	return max
}
