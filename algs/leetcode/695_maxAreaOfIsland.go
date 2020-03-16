package leetcode

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	visited := make([][]bool, row)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, col)
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 {
			return 0
		}

		grid[x][y] = 0 // 标记访问过的岛屿, 不会重复访问
		return 1 + dfs(x+1, y) + dfs(x-1, y) + dfs(x, y+1) + dfs(x, y-1)
	}

	max := -1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				if cur := dfs(i, j); cur > max {
					max = cur
				}
			}
		}
	}

	return max
}
