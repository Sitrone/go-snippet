package algs

type KMP struct {
	pattern string
	next    []int
}

// ref : https://www.nayuki.io/page/knuth-morris-pratt-string-matching
// 方法一：算法4 使用DFA构造next表
// 方法二：直接构造  https://www.zhihu.com/question/21923021 已经截图印象笔记

func NewKMP(pattern string) *KMP {
	next := computeNext(pattern)
	return &KMP{pattern: pattern, next: next}
}

func (kmp *KMP) Search(text string) int {
	i, j := 0, 0

	for i < len(text) && j < len(kmp.pattern) {
		if j == -1 || text[i] == kmp.pattern[j] {
			i++
			j++
		} else {
			j = kmp.next[j]
		}
	}

	if j == len(kmp.pattern) {
		return i - j
	} else {
		return -1
	}
}

func computeNext(pattern string) []int {
	next := make([]int, len(pattern))

	next[0] = -1
	i, j := 0, -1

	for i < len(pattern)-1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}
