package leetcode

import "strings"

func sortString(s string) string {
	if len(s) < 2 {
		return s
	}

	m := make(map[int32]int, 26)
	for _, c := range s {
		m[c] += 1
	}

	var buffer strings.Builder
	buffer.Grow(len(s))
	for len(m) > 0 {
		for i := 0; i < 26; i++ {
			if m[int32('a'+i)] > 0 {
				buffer.WriteByte(byte('a' + i))

				if m[int32('a'+i)]--; m[int32('a'+i)] == 0 {
					delete(m, int32('a'+i))
				}
			}
		}

		for i := 25; i >= 0; i-- {
			if m[int32('a'+i)] > 0 {
				buffer.WriteByte(byte('a' + i))
				if m[int32('a'+i)]--; m[int32('a'+i)] == 0 {
					delete(m, int32('a'+i))
				}
			}
		}
	}

	return buffer.String()
}

func sortString1(s string) string {
	if len(s) < 2 {
		return s
	}

	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}

	l := len(s)

	var buffer strings.Builder
	buffer.Grow(l)
	for buffer.Len() < l {
		for i := 0; i < 26 && buffer.Len() < l; i++ {
			if m[i] > 0 {
				buffer.WriteByte(byte('a' + i))
				m[i]--
			}
		}

		for i := 25; i >= 0 && buffer.Len() < l; i-- {
			if m[i] > 0 {
				buffer.WriteByte(byte('a' + i))
				m[i]--
			}
		}
	}

	return buffer.String()
}
