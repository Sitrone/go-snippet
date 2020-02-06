package leetcode

import (
	"fmt"
	"testing"
)

func TestAddbinary(t *testing.T) {
	var (
		a = "1010"
		b = "1011"
	)

	fmt.Println(addBinary(a, b))
}

func TestAddbinary1(t *testing.T) {
	var (
		a = "10001"
		b = "1111111"
	)

	fmt.Println(addBinary(a, b))
}

func TestAddbinary2(t *testing.T) {
	var (
		a = "1"
		b = "1111111"
	)

	fmt.Println(addBinary(a, b))
}
