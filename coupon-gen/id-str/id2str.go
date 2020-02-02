package id_str

import "math"

// 按照进制的思路，进行id <-> string 的相互转换

const (
	defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	defaultLen      = 6
)

type Option func(*IdStrData)

func NewIdStr(options ...Option) *IdStrData {
	d := &IdStrData{
		Alphabet: []rune(defaultAlphabet),
		Len:      defaultLen,
	}
	for _, o := range options {
		o(d)
	}

	return d
}

func WithAlphabet(alphabet string) Option {
	return func(d *IdStrData) {
		d.Alphabet = []rune(alphabet)
	}
}

func WithLen(len int) Option {
	return func(d *IdStrData) {
		d.Len = len
	}
}

type IdStrData struct {
	// Alphabet is the alphabet used to generate new ids
	Alphabet []rune

	//length of a generated id
	Len int
}

// 数字 -> 字符串映射, 思路：将给定字符串看成进制，这样就可以计算出第n个数对应的字符，高位不存在的用最低位补齐
// 引申： uid -> 字符串映射
func (c *IdStrData) Id2String(id int) string {
	alphabetLen := len(c.Alphabet)

	ret := make([]rune, c.Len)
	for i := 0; i < len(ret); i++ {
		ret[i] = c.Alphabet[0]
	}

	var (
		i   = c.Len - 1
		num int
	)

	for id != 0 {
		id, num = id/alphabetLen, id%alphabetLen
		ret[i] = c.Alphabet[num]
		i--
	}

	return string(ret)
}

func (c *IdStrData) Str2Id(s string) int {
	alphabetLen := len(defaultAlphabet)

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

func (c *IdStrData) MaxId() int {
	return int(math.Pow(float64(len(c.Alphabet)), float64(c.Len)))
}

func (c *IdStrData) indexOfAlphabet(r rune) int {
	for i, v := range c.Alphabet {
		if v == r {
			return i
		}
	}

	return 0
}
