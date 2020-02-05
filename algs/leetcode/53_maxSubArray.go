package leetcode

// Kadane算法扫描一次整个数列的所有数值，
// 在每一个扫描点计算以该点数值为结束点的子数列的最大和（正数和）。
// 该子数列由两部分组成：以前一个位置为结束点的最大子数列、该位置的数值。
// 因为该算法用到了“最佳子结构”（以每个位置为终点的最大子数列都是基于其前一位置的最大子数列计算得出,
// 该算法可看成动态规划的一个例子。
// TODO
// dp
// 状态转移方程：sum[i] = max{sum[i-1]+a[i],a[i]}
// [3,-1, 3]若[i-1]之前都连续，则 [i-1]+[i]也连续； sum[2] = 2，而不是3；sum数组是中间结果，不是答案
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxTwo(dp[i-1]+nums[i], nums[i])
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	last, cur := nums[0], 0

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = maxTwo(last+nums[i], nums[i])
		last = cur
		if cur > max {
			max = cur
		}
	}

	return max
}
