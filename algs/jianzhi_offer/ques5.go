package jianzhi_offer

import (
	"fmt"
	"strings"
)

// 从尾到头打印链表
func PrintListFromEnd2Front(head *ListNode) {
	if head == nil {
		return
	}

	next := head.Next
	PrintListFromEnd2Front(next)
	fmt.Println(head.Val)
}

func PrintListFromEnd2Front1(head *ListNode) {
	if head == nil {
		return
	}

	var stack Stack
	cur := head
	for cur != nil {
		stack.Push(cur.Val)
		cur = cur.Next
	}

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}

// 递归打印
func reversePrint(head *ListNode) []int {
	ans := make([]int, 0, 10)
	var f func(cur *ListNode)
	f = func(cur *ListNode) {
		if cur == nil {
			return
		}

		f(cur.Next)
		ans = append(ans, cur.Val)
	}

	f(head)
	return ans
}

// 普通打印，然后翻转
func reversePrint1(head *ListNode) []int {
	ans := make([]int, 0, 10)
	for cur := head; cur != nil; cur = cur.Next {
		ans = append(ans, cur.Val)
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return ans
}

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
