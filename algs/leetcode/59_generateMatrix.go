package leetcode

// 类似题目54
func generateMatrix(n int) [][]int {
	mat := make([][]int, n)

	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	var (
		l, r, t, b = 0, n - 1, 0, n - 1 // 四条边
		num, tar   = 1, n * n
	)
	for num <= tar {
		for i := l; i <= r; i++ {
			mat[t][i] = num // left to right.
			num++
		}
		t++
		for i := t; i <= b; i++ {
			mat[i][r] = num // top to bottom.
			num++
		}
		r--
		for i := r; i >= l; i-- {
			mat[b][i] = num // right to left.
			num++
		}
		b--
		for i := b; i >= t; i-- {
			mat[i][l] = num // bottom to top.
			num++
		}
		l++
	}
	return mat
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var (
		row = len(matrix)
		col = len(matrix[0])

		ret = make([]int, row*col)
	)

	var (
		l, r, t, b = 0, col - 1, 0, row - 1 // 四条边
		num, tar   = 0, row * col
	)
	for num < tar {
		for i := l; i <= r && num < tar; i++ {
			ret[num] = matrix[t][i] // left to right.
			num++
		}
		t++
		for i := t; i <= b && num < tar; i++ {
			ret[num] = matrix[i][r] // top to bottom.
			num++
		}
		r--
		for i := r; i >= l && num < tar; i-- {
			ret[num] = matrix[b][i] // right to left.
			num++
		}
		b--
		for i := b; i >= t && num < tar; i-- {
			ret[num] = matrix[i][l] // bottom to top.
			num++
		}
		l++
	}
	return ret
}
