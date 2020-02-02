package leetcode

func isValid(s string) bool {
	rs := []rune(s)
	if len(rs) == 0 {
		return true
	}
	if len(rs)&0x01 != 0 {
		return false
	}

	rightMap := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	var (
		last  rune
		stack []rune
	)
	stack = append(stack, rs[0])
	for i := 1; i < len(rs); i++ {
		if v, ok := rightMap[rs[i]]; ok {
			last = stack[len(stack)-1]
			if last != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, rs[i])
		}
	}

	return len(stack) == 0
}
