package leetcode

// 大数乘法
// 思路：拆解成num2的每一位和num1相乘的结果相加 时间复杂度：O(M N)。M,NM,N 分别为 num1 和 num2 的长度。 （PS：此为横式乘法）
// 用时4ms，击败70%用户，内存3.2MB，击败40%用户
// FIXME 优化方案，使用竖式乘法，https://leetcode-cn.com/problems/multiply-strings/solution/you-hua-ban-shu-shi-da-bai-994-by-breezean/
func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Bytes, num2Bytes := []byte(num1), []byte(num2)
	sumStack := multiplyOneDigit(num1Bytes, num2Bytes[len(num2Bytes)-1]-'0')

	var k = 1
	for i := len(num2Bytes) - 2; i >= 0; i, k = i-1, k+1 {
		//oneMul := append(multiplyOneDigitWithAmplify(num1Bytes, num2Bytes[i]-'0', k), fillZero(k)...)
		oneMul := multiplyOneDigitWithAmplify(num1Bytes, num2Bytes[i]-'0', k)
		sumStack = addTwoString(sumStack, oneMul)
	}

	return string(sumStack)
}

func fillZero(n int) []byte {
	ans := make([]byte, 0, n)
	for ; n > 0; n-- {
		ans = append(ans, '0')
	}
	return ans
}

func multiplyOneDigitWithAmplify(num []byte, mul uint8, times int) []byte {
	var (
		mulRet, canary uint8
		buf            = make([]byte, len(num)+1+times)
	)

	i, k := len(num)-1, 0

	for ; times > 0; times-- {
		buf[k] = '0'
		k++
	}

	for ; i >= 0; i, k = i-1, k+1 {
		mulRet = (num[i]-'0')*mul + canary
		mulRet, canary = mulRet%10, mulRet/10

		buf[k] = mulRet + '0'
	}

	if canary > 0 {
		buf[k] = canary + '0'
		k++
	}

	ans := buf[:k]
	for j := 0; j < len(ans)/2; j++ {
		ans[j], ans[len(ans)-j-1] = ans[len(ans)-j-1], ans[j]
	}

	return ans
}

func multiplyOneDigit(num []byte, mul uint8) []byte {
	var (
		mulRet, canary uint8
		buf            = make([]byte, len(num)+1)
	)

	i, k := len(num)-1, 0
	for ; i >= 0; i, k = i-1, k+1 {
		mulRet = (num[i]-'0')*mul + canary
		mulRet, canary = mulRet%10, mulRet/10

		buf[k] = mulRet + '0'
	}

	if canary > 0 {
		buf[k] = canary + '0'
		k++
	}

	ans := buf[:k]
	for j := 0; j < len(ans)/2; j++ {
		ans[j], ans[len(ans)-j-1] = ans[len(ans)-j-1], ans[j]
	}

	return ans
}

func addTwoString(a, b []byte) []byte {
	var (
		sum, canary uint8
		buf         []byte
	)

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

	return ans
}
