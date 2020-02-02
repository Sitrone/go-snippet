package coupon

import (
	"fmt"
	"testing"
)

func TestCoupon2Data_Gen(t *testing.T) {
	coupon := NewData2()
	coupon.Alphabet = []rune("023456789ABCDEFGHJKMNOPQRSTUVWXYZ")
	//coupon.Alphabet = []rune("0123456789")
	coupon.Len = 6
	ret, _ := coupon.Gen(500_000)

	fmt.Println(len(ret))
	//fmt.Println(ret)
}
