package coupon

import (
	"errors"
	"math"
)

type Coupon2Data struct {
	// Alphabet is the alphabet used to generate new ids
	Alphabet []rune

	//length of a generated id
	Len int
}

func NewData2() *Coupon2Data {
	return &Coupon2Data{
		Alphabet: []rune(defaultAlphabet),
		Len:      defaultLen,
	}
}

func (c *Coupon2Data) Gen(nums int) ([]string, error) {
	if len(c.Alphabet) < c.Len {
		return nil, errors.New("not enough alphabet to generate len unique string")
	}

	if nums > c.total() {
		nums = c.total()
	}

	ret := make([]string, nums)
	m := make(map[string]struct{}, nums)

	total := c.total()
	for i := 0; i < nums; {
		pos := srand.Int63n(int64(total))
		s := c.pos2String(int(pos))
		if _, ok := m[s]; !ok {
			ret[i] = s
			i++
		}
	}

	return ret, nil
}

func (c *Coupon2Data) pos2String(pos int) string {
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

func (c *Coupon2Data) total() int {
	return int(math.Pow(float64(len(c.Alphabet)), float64(c.Len)))
}
