package jianzhi_offer

// 思路1 hashMap
//type WordsFrequency struct {
//	m map[string]int
//}
//
//func Constructor(book []string) WordsFrequency {
//	f := WordsFrequency{m: make(map[string]int, 16)}
//	for _, s := range book {
//		f.m[s]++
//	}
//	return f
//}
//
//func (this *WordsFrequency) Get(word string) int {
//	return this.m[word]
//}

// 思路2trie tree
type WordsFrequency struct {
	root *Node
}

type Node struct {
	next  map[byte]*Node
	isEnd bool
	val   int
}

func (this *WordsFrequency) addWord(word string) {
	var n = this.root
	for i := 0; i < len(word); i++ {
		if n.next[word[i]] == nil {
			n.next[word[i]] = &Node{
				next: make(map[byte]*Node, 10),
			}
		}
		n = n.next[word[i]]
	}

	n.isEnd = true
	n.val += 1
}

func Constructor(book []string) WordsFrequency {
	f := WordsFrequency{
		root: &Node{
			next: make(map[byte]*Node, 26),
		},
	}
	for _, s := range book {
		f.addWord(s)
	}
	return f
}

func (this *WordsFrequency) Get(word string) int {
	for i, n := 0, this.root; i < len(word) && n.next[word[i]] != nil; i++ {
		n = n.next[word[i]]
		if n.isEnd && i == len(word)-1 {
			return n.val
		}
	}

	return 0
}
