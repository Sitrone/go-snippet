package leetcode

import "bytes"

func removeOuterParentheses(S string) string {
	var (
		ret   bytes.Buffer
		level int
	)

	for _, r := range []rune(S) {
		if r == ')' {
			level--
		}
		if level >= 1 {
			ret.WriteRune(r)
		}
		if r == '(' {
			level++
		}
	}
	return ret.String()
}
