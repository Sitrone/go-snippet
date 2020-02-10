package leetcode

// 双指针
func maxArea(height []int) int {
	left, right := 0, len(height)-1

	maxTwo := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	minTwo := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	var ans int
	for left < right {
		ans = maxTwo(ans, (right-left)*minTwo(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return ans
}
