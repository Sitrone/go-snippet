package leetcode

import "math"

// 找到最大的两个值
func maxProduct1(nums []int) int {
	var (
		max    int
		max2nd int
	)

	if nums[0] > nums[1] {
		max = nums[0]
		max2nd = nums[1]
	} else {
		max = nums[1]
		max2nd = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > max {
			max2nd, max = max, nums[i]
		} else if nums[i] > max2nd {
			max2nd = nums[i]
		}
	}

	return (max - 1) * (max2nd - 1)
}

// 暴力法
func maxProduct2(nums []int) int {
	var (
		max     = math.MinInt32
		calFunc func(x, y int) int
		maxFunc func(x, y int) int
	)

	calFunc = func(x, y int) int {
		return (x - 1) * (y - 1)
	}
	maxFunc = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			max = maxFunc(max, calFunc(nums[i], nums[j]))
		}
	}

	return max
}
