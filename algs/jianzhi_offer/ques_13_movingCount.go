package jianzhi_offer

func movingCount(m int, n int, k int) int {
	digitSum := func(x int) int {
		var total int
		for ; x > 0; x /= 10 {
			total += x % 10
		}

		return total
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || digitSum(x)+digitSum(y) > k || visited[x][y] {
			return 0
		}

		visited[x][y] = true
		return dfs(x+1, y) + dfs(x, y+1) + 1
	}

	return dfs(0, 0)
}
