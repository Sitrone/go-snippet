package leetcode

import "strings"

var index = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{}, len(words))
	buffer := strings.Builder{}
	for _, word := range words {
		buffer.Reset()
		for _, b := range []byte(word) {
			buffer.WriteString(index[b-'a'])
		}
		m[buffer.String()] = struct{}{}
	}

	return len(m)
}
