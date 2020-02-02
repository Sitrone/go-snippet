package leetcode

// dfs+回溯+剪枝
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0, k)

	var sumDfs func(k, reminder, start int)

	sumDfs = func(k, reminder, start int) {
		if len(cur) == k {
			if reminder == 0 {
				tmp := make([]int, k)
				copy(tmp, cur)
				ans = append(ans, tmp)
			}

			return
		}

		// 此处可以剪枝
		if reminder > 0 && len(cur) < k {
			for i := start + 1; i < 10; i++ {
				cur = append(cur, i)
				sumDfs(k, reminder-i, i)
				cur = cur[:len(cur)-1]
			}
		}
	}

	sumDfs(k, n, 0)
	return ans
}
