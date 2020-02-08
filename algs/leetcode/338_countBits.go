package leetcode

//  i & (i - 1）是比 i 小的，而且i & (i - 1)的1的个数已经在前面算过了，所以i的1的个数就是 i & (i - 1)的1的个数加上1
func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}

	return ans
}
