package coupon

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}

	for i := 0; i < 5; i++ {
		dest := make([]int, len(arr))
		copy(dest, arr)
		shuffle(dest)
		fmt.Println(dest)
	}
}
