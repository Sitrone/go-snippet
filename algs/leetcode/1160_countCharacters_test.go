package leetcode

import (
	"fmt"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	words := []string{"hello", "world", "leetcode"}
	chars := "welldonehoneyr"

	fmt.Println(countCharacters(words, chars))
}
