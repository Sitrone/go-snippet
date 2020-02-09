package leetcode

import (
	"strconv"
)

// 思路1,转换成字符串检查
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	var ans = true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			ans = false
			break
		}
	}

	return ans
}

// 思路2：直接判断
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0)
	var totalDigit int
	for tmp := x; tmp > 0; totalDigit++ {
		nums = append(nums, tmp%10)
		tmp /= 10
	}

	var ans = true
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			ans = false
			break
		}
	}

	return ans
}
