package jianzhi_offer

import "fmt"

// 超时
func lastRemaining(n int, m int) int {
	if n < 1 || m < 0 {
		return n
	}

	cycle := make([]int, n)
	for i := 0; i < n; i++ {
		cycle[i] = i
	}

	var pos int
	for len(cycle) != 1 {
		pos--
		for i := 0; i < m; i++ {
			pos %= len(cycle)
			pos++
		}

		pos %= len(cycle)
		fmt.Printf("%d ", cycle[pos])
		cycle = append(cycle[0:pos], cycle[pos+1:]...)
	}

	return cycle[0]
}

// TODO
func lastRemaining1(n int, m int) int {

	return 0
}
