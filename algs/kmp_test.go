package algs

import (
	"fmt"
	"testing"
)

func TestKMP_Search(t *testing.T) {
	kmp := NewKMP("abababca")

	// 6
	fmt.Println(kmp.Search("ababababca"))
}
