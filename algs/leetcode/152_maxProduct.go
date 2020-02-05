package leetcode

import "math"

// dp
// TODO 遍历数组时计算当前最大值，不断更新
// 由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值imin，imin = min(imin * nums[i], nums[i])
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	minTwo := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	imin, imax, max := 1, 1, math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		imax = maxTwo(imax*nums[i], nums[i])
		imin = minTwo(imin*nums[i], nums[i])

		max = maxTwo(imax, max)
	}

	return max
}
