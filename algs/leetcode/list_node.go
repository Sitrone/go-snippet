package leetcode

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var ret []int
	tmp := l
	for tmp != nil {
		ret = append(ret, tmp.Val)
		tmp = tmp.Next
	}

	return fmt.Sprintf("listNodes:%v", ret)
}
