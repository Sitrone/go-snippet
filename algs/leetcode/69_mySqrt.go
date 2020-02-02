package leetcode

// https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	start, end := 1, x
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid + 1
		} else if mid*mid > x {
			end = mid - 1
		}
	}

	return start - 1
}
