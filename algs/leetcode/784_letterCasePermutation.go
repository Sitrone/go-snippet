package leetcode

import "strings"

func letterCasePermutation(S string) []string {
	ans := make([]string, 0, 2)
	S = strings.ToLower(S)

	var backtrace func(pos int, cur []byte)
	backtrace = func(pos int, cur []byte) {
		if len(cur) == len(S) {
			ans = append(ans, string(cur))
			return
		}

		if S[pos] >= 'a' && S[pos] <= 'z' {
			for i := 0; i < 2; i++ {
				cur = append(cur, S[pos]-uint8(32*i))
				backtrace(pos+1, cur)
				cur = cur[:len(cur)-1]
			}
		} else {
			cur = append(cur, S[pos])
			backtrace(pos+1, cur)
		}
	}

	cur := make([]byte, 0, len(S))
	backtrace(0, cur)
	return ans
}
