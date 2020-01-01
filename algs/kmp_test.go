package algs

import (
	"fmt"
	"testing"
)

func TestKMP_Search(t *testing.T) {
	kmp := NewKMP("abababca")

	// 6
	fmt.Println(kmp.Search("ababababca"))

	var s = []int{1, 2, 3}
	r := fmt.Sprintf("%q\n", s)
	fmt.Println(r)

	var u = "hello,世界"
	us := []rune(u)
	for i, j := 0, len(us)-1; i < j; i, j = i+1, j-1 {
		us[i], us[j] = us[j], us[i]
	}
	fmt.Println(string(us))
}
