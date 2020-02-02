package snippet

import "math"

// https://time.geekbang.org/column/article/74287
// 回溯算法两个经典例子

// 0-1 背包
func Bag01(items []int, w int) int {
	var maxW = math.MinInt8

	var bag func(i, cw int, items []int, n, w int)

	// i, 考察到第i个物品， cw 当前物品的总总量
	// items，所有物品的重量，n 物品的总个数，w 背包最大承重
	bag = func(i, cw int, items []int, n, w int) {
		if i == n {
			if cw > maxW {
				maxW = cw
			}
			return
		}

		bag(i+1, cw, items, n, w) // 不放第i个
		if cw+items[i] < w {
			bag(i+1, cw+items[i], items, n, w) // 放入第i个
		}
	}

	bag(0, 0, items, len(items), w)
	return maxW
}

// 正则表达式匹配
// 假设正则表达式中只包含“*”和“?”这两种通配符，并且对这两个通配符的语义稍微做些改变，
// 其中，“*”匹配任意多个（大于等于 0 个）任意字符，“?”匹配零个或者一个任意字符
func RegMatch(text, pattern string) bool {
	rText := []rune(text)
	rPattern := []rune(pattern)
	tLen := len(rText)
	pLen := len(rPattern)

	var (
		match   func(tPos, pPos int, text, pattern []rune)
		matched bool
	)
	match = func(tPos, pPos int, text, pattern []rune) {
		if matched {
			return
		}

		if pPos == pLen && tPos == tLen {
			matched = true
			return
		}
		if pPos == pLen || tPos == tLen {
			return
		}

		if pattern[pPos] == '*' { // *匹配任意个字符
			for i := 0; i < tLen-tPos; i++ {
				match(tPos+i, pPos+1, text, pattern)
			}
		} else if pattern[pPos] == '?' { // ?匹配0个或者1个字符
			match(tPos, pPos+1, text, pattern)
			match(tPos+1, pPos+1, text, pattern)
		} else if text[tPos] == pattern[pPos] {
			match(tPos+1, pPos+1, text, pattern)
		}
	}

	match(0, 0, rText, rPattern)
	return matched
}
