package leetcode

import "sort"

func diagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 {
		return nil
	}

	minTwo := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	h, w := len(mat), len(mat[0])
	cur := make([]int, 0, minTwo(h, w))

	for j := 0; j < w; j++ {
		cur = cur[:0]
		for i, k := 0, j; i < h && k < w; i, k = i+1, k+1 {
			cur = append(cur, mat[i][k])
		}

		sort.Ints(cur)
		idx := 0
		for i, k := 0, j; i < h && k < w; i, k = i+1, k+1 {
			mat[i][k] = cur[idx]
			idx++
		}
	}

	for i := 1; i < h; i++ {
		cur = cur[:0]

		for j, k := 0, i; j < w && k < h; j, k = j+1, k+1 {
			cur = append(cur, mat[k][j])
		}

		sort.Ints(cur)
		idx := 0
		for j, k := 0, i; j < w && k < h; j, k = j+1, k+1 {
			mat[k][j] = cur[idx]
			idx++
		}
	}

	return mat
}
