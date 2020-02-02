package leetcode

// 即找到数据中第一个大于等于target的值
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var mid int

	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return low
}
