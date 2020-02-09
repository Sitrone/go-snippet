package leetcode

import (
	"fmt"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	S := "a1b2"
	ans := letterCasePermutation(S)

	fmt.Println(len(ans))
	fmt.Println(ans)
}
