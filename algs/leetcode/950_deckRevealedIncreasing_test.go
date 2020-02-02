package leetcode

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	var arr = []int{17, 13, 11, 2, 3, 5, 7}
	ret := deckRevealedIncreasing(arr)
	fmt.Println(ret)
}
