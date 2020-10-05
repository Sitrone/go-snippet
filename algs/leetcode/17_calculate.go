package leetcode

func calculate(s string) int {
	var (
		x, y  = 1, 0
		aFunc func(x, y int) int
		bFunc func(x, y int) int
	)
	aFunc = func(x, y int) int {
		return 2*x + y
	}

	bFunc = func(x, y int) int {
		return x + 2*y
	}

	for _, b := range s {
		if b == 'A' {
			x = aFunc(x, y)
		} else if b == 'B' {
			y = bFunc(x, y)
		}
	}

	return x + y
}

func calculate1(s string) int {
	return len(s) * 2
}
