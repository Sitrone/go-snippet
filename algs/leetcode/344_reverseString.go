package leetcode

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	for i, j := 0, len(s)-1; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
