package jianzhi_offer

// 数组旋转

// 魔术索引，没有说数组有序，无法用二分法
func findMagicIndex(nums []int) int {
	var ans = -1
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			ans = i
			break
		}
	}

	return ans
}
