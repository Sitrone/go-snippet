package leetcode

import (
	"fmt"
	"testing"
)

func TestCoinsChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
