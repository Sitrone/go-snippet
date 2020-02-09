package leetcode

func numberOfSteps(num int) int {
	var f func(step, reminder int) int
	f = func(step, reminder int) int {
		if reminder == 0 {
			return step
		}

		step++
		if (reminder & 0x01) == 0 {
			return f(step, reminder>>1)
		} else {
			return f(step, reminder-1)
		}
	}

	return f(0, num)
}
