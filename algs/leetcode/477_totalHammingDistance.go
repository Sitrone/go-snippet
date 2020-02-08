package leetcode

// 暴力解法超时
func totalHammingDistance(nums []int) int {
	hammingDistance := func(x int, y int) int {
		numberOf1 := func(n int) int {
			var total int
			for ; n > 0; total++ {
				n &= n - 1
			}

			return total
		}

		return numberOf1(x ^ y)
	}

	var total int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			total += hammingDistance(nums[i], nums[j])
		}
	}

	return total
}

// 参考答案
// 从逐一数字的横向比较，转变为纵向的逐位比较
// 对于每一个数的某一位来说，若当前位为 1，那么对于当前位来说，所有数字的同位上为 0 的个数即当前数字当前位的汉明距离
func totalHammingDistance1(nums []int) int {
	numLen := len(nums)
	sum := 0
	for j := 0; j < 32; j++ {
		ones := 0
		for i := 0; i < numLen; i++ {
			ones += (nums[i] >> uint(j)) & 1
		}
		sum += ones * (numLen - ones)
	}
	return sum
}
