package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) < 10 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}

		return -1
	}

	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
