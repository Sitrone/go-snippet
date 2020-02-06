package leetcode

func addToArrayForm(A []int, K int) []int {
	var (
		sum, canary, i, k, b int
		ans                  []int
	)

	lenK := 0
	for tmp := K; tmp > 0; tmp /= 10 {
		lenK++
	}
	if len(A) > lenK {
		ans = make([]int, len(A)+1)
	} else {
		ans = make([]int, lenK+1)
	}

	for i = len(A) - 1; i >= 0 || K > 0; i, k = i-1, k+1 {
		sum = canary
		if K > 0 {
			K, b = K/10, K%10
			sum += b
		}
		if i >= 0 {
			sum += A[i]
		}

		sum, canary = sum%10, sum/10
		ans[k] = sum
	}

	if canary > 0 {
		ans[k] = canary
		k++
	}

	ans = ans[:k]
	for j := 0; j < len(ans)/2; j++ {
		ans[j], ans[len(ans)-j-1] = ans[len(ans)-j-1], ans[j]
	}

	return ans
}
