package jianzhi_offer

func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	return string(append(bytes[n:], bytes[:n]...))
}
