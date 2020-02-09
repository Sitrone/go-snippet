package leetcode

import (
	"bytes"
	"fmt"
)

type node struct {
	next  map[uint8]*node
	isEnd bool
}

func (n *node) String() string {
	return fmt.Sprintf("isEnd=%t, next=%v", n.isEnd, n.next)
}

func addWord(n *node, word string) {
	for i := 0; i < len(word); i++ {
		if n.next[word[i]-'a'] == nil {
			n.next[word[i]-'a'] = &node{next: make(map[uint8]*node, 26)}
		}
		n = n.next[word[i]-'a']
	}
	n.isEnd = true
}

// 思路1：使用Trie树，4ms，击败6.37%用户，性能损耗在需要先扫描全部字符串，然后再判断
// 占用存储4.7MB，击败5.1%用户
func longestCommonPrefix(strs []string) string {
	var root = &node{next: make(map[uint8]*node, 26)}

	for _, s := range strs {
		if s == "" {
			return ""
		}
		addWord(root, s)
	}

	fmt.Println(root)
	var ans bytes.Buffer
	for cur := root; cur != nil && len(cur.next) == 1 && cur.isEnd == false; {
		for k, n := range cur.next {
			ans.WriteByte(k + 'a')
			cur = n
		}
	}

	return ans.String()
}

// 直接横向扫描, 耗时0ms，击败100%， 存储2.4MB，击败98.56%
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var ans = make([]byte, 0, 2)
	s1 := strs[0]
	var equal bool
	for i := 0; i < len(s1); i++ {
		equal = true
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != s1[i] {
				equal = false
			}
		}
		if equal {
			ans = append(ans, s1[i])
		} else {
			break
		}
	}

	return string(ans)
}
