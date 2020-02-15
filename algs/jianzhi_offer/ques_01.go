package jianzhi_offer

func CheckPermutation(s1 string, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)

	m := make(map[rune]int, len(r1)/2)
	for _, r := range r1 {
		m[r]++
	}

	for _, r := range r2 {
		if _, ok := m[r]; ok {
			m[r]--
		} else {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func isUnique(astr string) bool {
	r := []rune(astr)
	m := make(map[rune]struct{}, len(r)/2)

	for _, v := range r {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}
