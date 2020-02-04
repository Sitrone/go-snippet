package leetcode

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var s = "9223372036854775808"

	fmt.Println(myAtoi1(s))
}
