package leetcode

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	n := len(matrix) - 1

	for i := 0; i < n; i++ {
		for j := i; j < n-i; j++ {
			matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]
		}
	}
}
