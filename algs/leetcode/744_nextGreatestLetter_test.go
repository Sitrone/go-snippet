package leetcode

import (
	"fmt"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	var bytes = []byte{'c', 'f', 'j'}
	var target = byte('g')
	fmt.Printf("%c\n", nextGreatestLetter(bytes, target))
}
