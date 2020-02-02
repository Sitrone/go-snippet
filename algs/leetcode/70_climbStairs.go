package leetcode

//https://leetcode-cn.com/problems/climbing-stairs/

// dp f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	ans := make([]int, n+1)
	ans[0], ans[1] = 1, 1

	for i := 2; i <= n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}

	return ans[n]
}
