package jianzhi_offer

func sumNums(n int) int {
	var f func(ret *int, n int) bool
	f = func(ret *int, n int) bool {
		*ret += n

		return n != 0 && f(ret, n-1)
	}

	var ans int
	f(&ans, n)

	return ans
}
