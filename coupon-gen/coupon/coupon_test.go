package coupon

import (
	"fmt"
	"testing"
)

func TestPermutationWithLen(t *testing.T) {
	//s := "023456789ABCDEFGHJKMNOPQRSTUVWXYZ"
	s := "0123456789"
	ret := PermutationWithLen(s, 6)
	//fmt.Println(ret)
	fmt.Println(len(ret))
}
