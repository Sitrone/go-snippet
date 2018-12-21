package algs

type KMP struct {
	pattern string
	lsp     []int
}

// ref : https://www.nayuki.io/page/knuth-morris-pratt-string-matching
// 算法4 使用DFA构造next表

func NewKMP(pattern string) *KMP {
	lsp := computeLspTable(pattern)
	return &KMP{pattern: pattern, lsp: lsp}
}

func (kmp *KMP) Search(text string) int {
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != kmp.pattern[j] {
			j = kmp.lsp[j-1]
		}
		if text[i] == kmp.pattern[j] {
			j++
			if j == len(kmp.pattern) {
				return i - (j - 1)
			}
		}
	}
	return -1
}

func computeLspTable(pattern string) []int {
	lsp := make([]int, len(pattern))
	lsp[0] = 0

	for i := 1; i < len(pattern); i++ {
		j := lsp[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lsp[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		lsp[i] = j
	}
	return lsp
}
