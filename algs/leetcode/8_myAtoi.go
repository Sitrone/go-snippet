package leetcode

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func myAtoi(str string) int {
	rs := []rune(str)

	// skip whitespace
	var i int
	for i = 0; i < len(rs); i++ {
		if rs[i] != ' ' {
			break
		}
	}
	rs = rs[i:]

	if len(rs) == 0 {
		return 0
	}

	var isNegative bool
	if rs[0] == '-' {
		isNegative = true
		rs = rs[1:]
	} else if rs[0] == '+' {
		rs = rs[1:]
	}

	var buf bytes.Buffer
	for j := 0; j < len(rs) && rs[j] >= '0' && rs[j] <= '9'; j++ {
		buf.WriteString(fmt.Sprintf("%c", rs[j]))
	}

	var ans int
	if buf.Len() > 0 {
		tmp, _ := strconv.ParseInt(buf.String(), 10, 64)
		if isNegative {
			tmp = -tmp
		}
		if tmp < math.MinInt32 {
			ans = math.MinInt32
		} else if tmp > math.MaxInt32 {
			ans = math.MaxInt32
		} else {
			ans = int(tmp)
		}
	}

	return ans
}

func myAtoi1(str string) int {
	rs := []rune(str)

	// skip whitespace
	var i int
	for i = 0; i < len(rs); i++ {
		if rs[i] != ' ' {
			break
		}
	}
	rs = rs[i:]

	if len(rs) == 0 {
		return 0
	}

	var isNegative bool
	if rs[0] == '-' {
		isNegative = true
		rs = rs[1:]
	} else if rs[0] == '+' {
		rs = rs[1:]
	}

	// skip first zero
	i = 0
	for ; i < len(rs) && rs[i] == '0'; i++ {
	}
	rs = rs[i:]

	var buf []rune
	for j := 0; j < len(rs) && rs[j] >= '0' && rs[j] <= '9'; j++ {
		buf = append(buf, rs[j])
	}

	pow10 := func(n int) int64 {
		var ret int64 = 1
		for k := 0; k < n; k++ {
			ret *= 10
		}
		return ret
	}

	var ans int
	if len(buf) > 0 {
		if len(buf) > 12 {
			if isNegative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		var tmp int64
		for j := len(buf) - 1; j >= 0; j-- {
			if isNegative {
				tmp -= int64(buf[j]-'0') * pow10(len(buf)-j-1)
				if tmp < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				tmp += int64(buf[j]-'0') * pow10(len(buf)-j-1)
				if tmp > math.MaxInt32 {
					return math.MaxInt32
				}
			}
		}

		ans = int(tmp)
	}

	return ans
}
