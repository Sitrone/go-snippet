package leetcode

// 4ms
func countPrimeSetBits(L int, R int) int {
	if L > R || L < 0 || R > 1000000 {
		return 0
	}

	primeBits := 0b10100010100010101100
	numberOf1 := func(n int) int {
		var total int
		for ; n > 0; total++ {
			n &= n - 1
		}

		return total
	}

	primeNum := 0
	for i := L; i <= R; i++ {
		primeNum += (primeBits >> numberOf1(i)) & 1
	}
	return primeNum
}

// 12ms 最简单直接的方法
func countPrimeSetBits1(L int, R int) int {
	if L > R || L < 0 || R > 1000000 {
		return 0
	}

	primeMap := map[int]bool{
		2: true, 3: true, 5: true, 7: true,
		11: true, 13: true, 17: true, 19: true,
	}

	numberOf1 := func(n int) int {
		var total int
		for ; n > 0; total++ {
			n &= n - 1
		}

		return total
	}

	primeNum := 0
	for i := L; i <= R; i++ {
		if primeMap[numberOf1(i)] {
			primeNum++
		}
	}
	return primeNum
}

//56ms 过度优化
func countPrimeSetBits2(L int, R int) int {
	if L > R || L < 0 || R > 1000000 {
		return 0
	}

	primeMap := map[int]bool{
		2: true, 3: true, 5: true, 7: true,
		11: true, 13: true, 17: true, 19: true,
	}

	numberOf1 := func(n int) int {
		var total int
		for ; n > 0; total++ {
			n &= n - 1
		}

		return total
	}

	primeNum := 0
	m := make(map[int]int, (R-L)/2)
	for i := L; i <= R; i++ {
		if v, ok := m[i&(i-1)]; ok {
			m[i] = v + 1
			if primeMap[v+1] {
				primeNum++
			}
		} else {
			of1 := numberOf1(i)
			m[i] = of1
			if primeMap[of1] {
				primeNum++
			}
		}

	}
	return primeNum
}
