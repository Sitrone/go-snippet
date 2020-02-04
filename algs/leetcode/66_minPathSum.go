package leetcode

// 思路1：回溯
// 思路2：dp dp[i][j] = min(dp[i-1][j], dp[i][j-1])+a[i][j]

// 二维数组
// 执行用时 : 8 ms , 在所有 Go 提交中击败了 91.75% 的用户
// 内存消耗 : 4.4 MB , 在所有 Go 提交中击败了 42.02% 的用户
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}

	row, col := len(grid), len(grid[0])
	var dp = make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < col; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[row-1][col-1]
}

func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}

	row, col := len(grid), len(grid[0])
	var dp = make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if i > 0 && j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[row-1][col-1]
}

// dp 一维数组 TODO
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	//min := func(x, y int) int {
	//	if x > y {
	//		return y
	//	}
	//
	//	return x
	//}

	row, col := len(grid), len(grid[0])
	var dp = make([][]int, row)

	return dp[row-1][col-1]
}
