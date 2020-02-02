package leetcode

import (
	"fmt"
	"testing"
)

func TestGroupPeople(t *testing.T) {
	var arr = []int{3, 3, 3, 3, 3, 1, 3}
	ret := groupThePeople2(arr)
	fmt.Println(ret)
}
