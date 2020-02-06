package leetcode

// 思路1：map
// 思路2：原地置换
func findDuplicate(nums []int) int {
	tmp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if tmp[nums[i]] != 0 {
			return nums[i]
		} else {
			tmp[nums[i]] = 1
		}
	}

	return -1
}

func findDuplicate1(nums []int) int {
	n := len(nums) - 1
	total := (1 + n) * n / 2

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum - total
}
