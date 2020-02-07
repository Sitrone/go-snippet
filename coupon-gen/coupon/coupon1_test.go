package coupon

import (
	"fmt"
	"log"
	"testing"
)

func TestCouponData_Gen(t *testing.T) {
	coupon := NewData()
	coupon.Alphabet = []rune("023456789ABCDEFGHJKMNOPQRSTUVWXYZ")
	//coupon.Alphabet = []rune("0123456789")
	coupon.Len = 6
	ret, _ := coupon.Gen(100)

	fmt.Println(len(ret))
	//fmt.Println(ret)
	m := make(map[string]struct{}, len(ret))
	for v := range ret {
		if _, ok := m[v]; ok {
			fmt.Println(v)
			log.Fatalf("duplicate key found, v=[%s]\n", v)
		}

		m[v] = struct{}{}
	}

}

func TestNewData(t *testing.T) {
	coupon := NewData()
	//coupon.Alphabet = []rune("023456789ABCDEFGHJKMNOPQRSTUVWXYZ")
	coupon.Alphabet = []rune("ABCDEFGHJKMNOPQRSTUVWXYZ")
	coupon.Len = 6

	str := coupon.pos2String(54032)
	fmt.Println(str)

	pos := coupon.string2Pos(str)
	fmt.Println(pos)
}
