package coupon

import (
	"errors"
	"math"
)

const (
	defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	defaultLen      = 6
)

type CouponData struct {
	// Alphabet is the alphabet used to generate new ids
	Alphabet []rune

	//length of a generated id
	Len int
}

type ICoupon interface {
	Gen(nums int) ([]string, error)
}

func NewData() *CouponData {
	return &CouponData{
		Alphabet: []rune(defaultAlphabet),
		Len:      defaultLen,
	}
}

func (c *CouponData) Gen(nums int) ([]string, error) {
	if len(c.Alphabet) < c.Len {
		return nil, errors.New("not enough alphabet to generate len unique string")
	}

	if nums > c.total() {
		nums = c.total()
	}

	return c.doGen(nums)
}

func (c *CouponData) doGen(nums int) ([]string, error) {
	arr := rnd.Perm(c.total())

	dest := make([]int, nums)
	for i := 0; i < nums; i++ {
		dest[i] = arr[i]
	}

	return c.map2Ret(dest), nil
}

func (c *CouponData) map2Ret(positions []int) []string {
	ret := make([]string, len(positions))

	for i, pos := range positions {
		ret[i] = c.pos2String(pos)
	}

	return ret
}

// 数字 -> 字符串映射, 思路：将给定字符串看成进制，这样就可以计算出第n个数对应的字符，高位不存在的用最低位补齐
// 引申： uid -> 字符串映射
func (c *CouponData) pos2String(pos int) string {
	alphabetLen := len(c.Alphabet)

	ret := make([]rune, c.Len)
	for i := 0; i < len(ret); i++ {
		ret[i] = c.Alphabet[0]
	}

	var (
		i     = c.Len - 1
		index int
	)

	for pos != 0 {
		pos, index = pos/alphabetLen, pos%alphabetLen
		ret[i] = c.Alphabet[index]
		i--
	}

	return string(ret)
}

func (c *CouponData) string2Pos(s string) int {
	alphabetLen := len(c.Alphabet)

	runes := []rune(s)

	var (
		ret        int
		beginCount bool
	)
	for i := 0; i < len(runes); i++ {
		if c.Alphabet[0] != runes[i] {
			beginCount = true
		}

		if beginCount {
			ret += int(math.Pow(float64(alphabetLen), float64(len(runes)-1-i))) * c.indexOfAlphabet(runes[i])
		}
	}

	return ret
}

func (c *CouponData) indexOfAlphabet(r rune) int {
	for i, v := range c.Alphabet {
		if v == r {
			return i
		}
	}

	return 0
}

func (c *CouponData) randomArr() []int {
	total := c.total()
	arr := make([]int, total, total)
	for i := 1; i < total; i++ {
		arr[i-1] = i
	}

	return arr
}

// 可能计算结果超亿，放入内存的话数据过大，需要优化
func (c *CouponData) total() int {
	return int(math.Pow(float64(len(c.Alphabet)), float64(c.Len)))
}
