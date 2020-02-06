package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoString(t *testing.T) {
	var (
		a = "111"
		b = "9991111111"
	)

	fmt.Println(string(addTwoString([]byte(a), []byte(b))))
}

func TestMultiplyOneDigit(t *testing.T) {
	var (
		nums       = "998"
		mul  uint8 = 9
	)

	fmt.Println(string(multiplyOneDigit([]byte(nums), mul)))
}

func TestMultiply(t *testing.T) {
	var (
		a = "123"
		b = "789"
	)

	fmt.Println(multiply(a, b))
}
