package leetcode

func oddCells(n int, m int, indices [][]int) int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}

	var (
		row, col int
	)
	for _, index := range indices {
		row, col = index[0], index[1]

		for i := 0; i < m; i++ {
			mat[row][i] += 1
		}

		for i := 0; i < n; i++ {
			mat[i][col] += 1
		}
	}

	var ans int
	for i, v := range mat {
		for j := range v {
			if mat[i][j]&0x01 != 0 {
				ans++
			}
		}
	}

	return ans
}
