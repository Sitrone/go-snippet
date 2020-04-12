package leetcode

import "sort"

func hasGroupsSizeX(deck []int) bool {
	var l = len(deck)
	if l < 1 {
		return false
	}

	sort.Ints(deck)

	for i := 2; i <= l/2; i++ {
		if l%i == 0 {
			has := true
			for j := 0; j < l; j += i {
				for k := j; k < i-1; k++ {
					if deck[k] != deck[k+1] {
						has = false
						goto OUT
					}
				}
			}
		OUT:
			if has {
				return true
			}
		}
	}

	return false
}

func hasGroupsSizeX1(deck []int) bool {
	var l = len(deck)
	if l < 1 {
		return false
	}

	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[deck[i]]++
	}

	if len(m) == 1 {
		if _, ok := m[1]; ok {
			return true
		}
	}

	for i := 2; i <= l/2; i++ {
		has := true
		for _, v := range m {
			if v%i != 0 {
				has = false
				break
			}
		}

		if has {
			return true
		}
	}

	return false
}
