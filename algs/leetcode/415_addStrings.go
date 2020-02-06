package leetcode

func addStrings(num1 string, num2 string) string {
	var (
		sum, canary uint8
		buf         []byte
	)

	a, b := []byte(num1), []byte(num2)

	pos1, pos2 := len(a)-1, len(b)-1

	if len(a) > len(b) {
		buf = make([]byte, len(a)+1)
	} else {
		buf = make([]byte, len(b)+1)
	}

	var k int
	for ; pos1 >= 0 || pos2 >= 0; k++ {
		sum = canary
		if pos1 >= 0 {
			sum += a[pos1] - '0'
			pos1--
		}
		if pos2 >= 0 {
			sum += b[pos2] - '0'
			pos2--
		}
		sum, canary = sum%10, sum/10
		buf[k] = sum + '0'
	}

	if canary > 0 {
		buf[k] = canary + '0'
		k++
	}

	ans := buf[:k]
	for j := 0; j < len(ans)/2; j++ {
		ans[j], ans[len(ans)-j-1] = ans[len(ans)-j-1], ans[j]
	}

	return string(ans)
}
