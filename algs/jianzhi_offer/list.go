package jianzhi_offer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	return fmt.Sprintf("ListNode{value=%d, next=%v}", head.Val, head.Next)
}

func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}

	pre := head
	cur := pre.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	head.Next = nil
	return pre
}

func NewList(elems ...int) *ListNode {
	if elems == nil {
		return nil
	}

	head := &ListNode{Val: elems[0]}
	cur := head
	for i := 1; i < len(elems); i++ {
		cur.Next = &ListNode{Val: elems[i]}
		cur = cur.Next
	}
	return head
}
