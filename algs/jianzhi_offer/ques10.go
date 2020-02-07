package jianzhi_offer

// n & (n-1) 去掉最右边的为1的二进制位
// 计算二进制中1的个数
func NumberOf1(n int) int {
	var total int
	for ; n > 0; total++ {
		n &= n - 1
	}

	return total
}

// 一条语句判断整数是不是2的整数次方
func IsPowOf2(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

// 计算异或去掉相同的位，然后计算1的个数
func Bit2Change(m, n int) int {
	return NumberOf1(m ^ n)
}
