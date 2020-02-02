package leetcode

// https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/

// 寻找比目标字母大的最小字母
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[0] > target || letters[len(letters)-1] <= target {
		return letters[0]
	}

	low, high := 0, len(letters)-1

	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if letters[mid] <= target {
			if mid == len(letters)-1 {
				return letters[len(letters)-1]
			}
			if letters[mid+1] > target {
				return letters[mid+1]
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}

	return ' '
}
