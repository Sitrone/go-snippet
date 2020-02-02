package leetcode

func subtractProductAndSum(n int) int {
	var (
		reminder, sum int
		mul           = 1
	)
	for n != 0 {
		n, reminder = n/10, n%10
		mul *= reminder
		sum += reminder
	}

	return mul - sum
}
