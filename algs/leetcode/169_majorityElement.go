package leetcode

import "sort"

// 求众数
// 思路1，排序取中间值 时间复杂度O(nLogn) 空间复杂度O(1)
// 思路2，利用快排求第k大的值
// 思路3，map，时间复杂度O(n), 空间复杂度O(n)
func majorityElement1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := nums[0]
	k := 1
	for i := 1; i < len(nums); i++ {
		if k == 0 {
			cur = nums[i]
		}

		if cur == nums[i] {
			k++
		} else {
			k--
		}
	}

	return cur
}

func majorityElement3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var m = make(map[int]int, 5)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return 0
}
