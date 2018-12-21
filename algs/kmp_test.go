package algs

import (
	"fmt"
	"testing"
)

func TestKMP_Search(t *testing.T) {
	kmp := NewKMP("aaab")

	// 6
	fmt.Println(kmp.Search("aaaaaaaaab"))
}
