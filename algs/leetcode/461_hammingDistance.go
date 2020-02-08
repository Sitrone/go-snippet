package leetcode

func hammingDistance(x int, y int) int {
	numberOf1 := func(n int) int {
		var total int
		for ; n > 0; total++ {
			n &= n - 1
		}

		return total
	}

	return numberOf1(x ^ y)
}
