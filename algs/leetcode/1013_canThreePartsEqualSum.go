package leetcode

// 双指针
func canThreePartsEqualSum(A []int) bool {
	var sum int
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	if sum%3 != 0 {
		return false
	}

	subSum := sum / 3
	var curSum, n int
	for i := 0; i < len(A); i++ {
		curSum += A[i]
		if curSum == subSum {
			n++
			curSum = 0
		}
	}
	return n > 2
}
