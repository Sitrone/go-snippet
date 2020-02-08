package leetcode

import "strings"

func letterCombinations(digits string) []string {
	if strings.TrimSpace(digits) == "" {
		return nil
	}

	for _, v := range digits {
		if v > '9' || v < '2' {
			return nil
		}
	}

	lettersMap := map[byte][]byte{
		'1': []byte("*"), '2': []byte("abc"), '3': []byte("def"),
		'4': []byte("ghi"), '5': []byte("jkl"), '6': []byte("mno"),
		'7': []byte("pqrs"), '8': []byte("tuv"), '9': []byte("wxyz"),
	}

	var (
		ans       = make([]string, 0)
		backtrace func(start int, curPath []byte)
	)

	backtrace = func(idx int, curPath []byte) {
		if idx == len(digits) {
			ans = append(ans, string(curPath))
			return
		}

		// 相当于穷举，一个一个的往前推进，在每一种数字里面有多个情况选或者不选
		for j := 0; j < len(lettersMap[digits[idx]]); j++ {
			curPath = append(curPath, lettersMap[digits[idx]][j])
			backtrace(idx+1, curPath)
			curPath = curPath[:len(curPath)-1]
		}
	}

	path := make([]byte, 0, len(digits))
	backtrace(0, path)
	return ans
}
