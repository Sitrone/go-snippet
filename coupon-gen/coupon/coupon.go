package coupon

import "math"

// 给定字符生成定长唯一字符串生成思路

// 思路0: hashids, 安全性一般
// 思路1: 枚举全部的指定长度的字符串，然后采用洗牌算法洗牌，最后取出需要大小的前面n条数据，枚举全部字符串, 安全性高,依赖于安全的伪随机数
// 思路2: 先生成给定字符能够组成所有字符串大小的一个数组，然后洗牌，取出前面需要的n条数据，从数据反算出来字符串，参考coupon1.go, 安全性高,依赖于安全的伪随机数
// 思路3: 每次取[0, 全部枚举字符大小]随机数, 然后反算出字符串，通过map来判重, 性能最高, 占用内存最小, 安全性高,依赖于安全的伪随机数
func PermutationWithLen(s string, l int) []string {
	if l > len(s) {
		return nil
	}

	runes := []byte(s)
	return doPermutationWithLen(runes, l, runes[0], runes[len(runes)-1])
}

func doPermutationWithLen(s []byte, l int, low, high byte) []string {
	ret := make([]string, 0, int(math.Pow(float64(len(s)), float64(l))))

	//// 优化成map, 测试变慢，优化算法
	//sourceMap := make(map[byte]int, len(s))
	//for i, b := range s {
	//	sourceMap[b] = i
	//}
	start, end := repeatRune(low, l), repeatRune(high, l)
	for !equalRunes(start, end) {
		ret = append(ret, string(start))
		incrRunes(start, s, low, high)
	}

	ret = append(ret, string(end))
	return ret
}

func incrRunes(s, source []byte, low, high byte) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == high {
			s[i] = low
		} else {
			s[i] = source[indexOfRunes(s[i], source)+1]
			//s[i] = source[sourceMap[s[i]]+1]
			break
		}
	}
}

func indexOfRunes(r byte, s []byte) int {
	for i := 0; i < len(s); i++ {
		if r == s[i] {
			return i
		}
	}

	return 0
}

func equalRunes(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func repeatRune(r byte, n int) []byte {
	ret := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		ret = append(ret, r)
	}

	return ret
}
