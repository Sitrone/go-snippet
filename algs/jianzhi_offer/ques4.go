package jianzhi_offer

import "strings"

// 替换空格为%20
func ReplaceBlank(s string) string {
	if s == "" {
		return ""
	}

	return strings.Replace(s, " ", "%20", -1)
}

func ReplaceBlank1(s string) string {
	ws := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ws++
		}
	}

	var ret strings.Builder
	ret.Grow(len(s) + ws*2)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret.WriteString("%20")
		} else {
			ret.WriteByte(s[i])
		}
	}
	return ret.String()
}
