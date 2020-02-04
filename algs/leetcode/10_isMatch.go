package leetcode

import "regexp"

// 正则表达式匹配
// 思路1： 调用stdku
// 思路2： 回溯
// 思路3： dp

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素

// 调用库函数
func isMatch(text, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return len(r.FindString(text)) == len(text) && r.MatchString(text)
}

// 回溯
func isMatch1(text string, pattern string) bool {
	var match func(text, pattern []rune) bool

	match = func(text, pattern []rune) bool {
		if len(pattern) == 0 {
			return len(text) == 0
		}

		var firstMatch bool
		if len(text) > 0 && (text[0] == pattern[0] || pattern[0] == '.') {
			firstMatch = true
		}

		if len(pattern) > 1 && pattern[1] == '*' {
			return match(text, pattern[2:]) || (firstMatch && match(text[1:], pattern))
		} else {
			return firstMatch && match(text[1:], pattern[1:])
		}
	}

	return match([]rune(text), []rune(pattern))
}

// dp 难点: 匹配*的状态转移方程
// ref：https://leetcode-cn.com/problems/regular-expression-matching/solution/dong-tai-gui-hua-zen-yao-cong-0kai-shi-si-kao-da-b/
// TODO
func isMatch2(text, pattern string) bool {

	return false
}
