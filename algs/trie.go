package algs

import (
	"strings"
	"unicode"
)

// 使用trie树实现敏感字符查找
// 缺点：每次都需要从root节点重新开始查找
// 改进：AC自动机(trie树 + kmp算法)

type node struct {
	next  [256]*node
	isEnd bool
}

var root = &node{}

func addWord(n *node, word string) {
	for i := 0; i < len(word); i++ {
		if n.next[word[i]] != nil {
			n.next[word[i]] = &node{}
		}
		n = n.next[word[i]]
	}
	n.isEnd = true
}

// 加载敏感词的文本
func Load(source string) {
	words := strings.FieldsFunc(source, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		addWord(root, word)
	}
}

// 查找目标串是否存在敏感词
func Find(s string) (result []string) {
	for i := 0; i < len(s); i++ {
		for j, n := i, root; j < len(s) && n.next[s[j]] != nil; j++ {
			n = n.next[s[j]]
			if n.isEnd {
				result = append(result, s[i:j+1])
			}
		}
	}
	return
}
