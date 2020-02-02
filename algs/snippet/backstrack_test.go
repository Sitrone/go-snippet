package snippet

import (
	"fmt"
	"testing"
)

func TestBag01(t *testing.T) {

}

func TestRegMatch(t *testing.T) {
	var (
		text    = "abcdddddf"
		pattern = "abc*f"
	)

	matched := RegMatch(text, pattern)
	fmt.Println(matched)
}

func TestRegMatch2(t *testing.T) {
	var (
		text    = "abcdddddf"
		pattern = "abc*?"
	)

	matched := RegMatch(text, pattern)
	fmt.Println(matched)
}
