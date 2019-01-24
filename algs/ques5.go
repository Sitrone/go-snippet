package algs

import "fmt"

// 从尾到头打印链表
func PrintListFromEnd2Front(head *ListNode) {
	if head == nil {
		return
	}

	next := head.Next
	PrintListFromEnd2Front(next)
	fmt.Println(head.Value)
}

func PrintListFromEnd2Front1(head *ListNode) {
	if head == nil {
		return
	}

	var stack Stack
	cur := head
	for cur != nil{
		stack.Push(cur.Value)
		cur = cur.Next
	}

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}