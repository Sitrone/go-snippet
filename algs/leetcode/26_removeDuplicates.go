package leetcode

// 思路 双指针
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}

	return idx + 1
}
