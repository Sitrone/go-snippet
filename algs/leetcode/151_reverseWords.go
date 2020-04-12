package leetcode

import "strings"

func reverseWords(s string) string {
	if len(strings.TrimSpace(s)) == 0 {
		return ""
	}

	fields := strings.Fields(s)
	var builder strings.Builder
	builder.Grow(len(s))
	for i := len(fields) - 1; i >= 0; i-- {
		builder.WriteString(fields[i])
		if i != 0 {
			builder.WriteByte(' ')
		}
	}

	return builder.String()
}
