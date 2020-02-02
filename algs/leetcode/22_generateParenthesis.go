package leetcode

// dfs
func generateParenthesis(n int) []string {
	var (
		ans = make([]string, 0)
		cur = make([]byte, 0, n*2)
	)

	parenthesisDfs(n, cur, &ans)
	return ans
}

func parenthesisDfs(n int, cur []byte, ans *[]string) {
	if len(cur) == n*2 {
		// cur -> string 放在判断成功之后，减少内存分配，提高性能
		if isLegalParenthesis(cur) {
			*ans = append(*ans, string(cur))
		}
		return
	}

	for _, b := range []byte{'(', ')'} {
		if getParenthesisLevel(cur) >= 0 { // 剪枝
			cur = append(cur, b)
			parenthesisDfs(n, cur, ans)
			cur = cur[:len(cur)-1]
		}
	}
}

func getParenthesisLevel(s []byte) int {
	var level int
	for _, v := range s {
		if level < 0 {
			return level
		}
		if v == '(' {
			level++
		} else {
			level--
		}
	}

	return level
}

func isLegalParenthesis(s []byte) bool {
	return getParenthesisLevel(s) == 0
}
