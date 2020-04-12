package leetcode

func numTilePossibilities(tiles string) int {
	m := make([]int, 26)
	for _, c := range tiles {
		m[c-'A']++
	}

	var dfs func() int
	dfs = func() int {
		var total int
		for i := 0; i < len(m); i++ {
			if m[i] > 0 {
				total++

				m[i]--
				total += dfs()
				m[i]++
			}
		}
		return total
	}

	return dfs()
}
