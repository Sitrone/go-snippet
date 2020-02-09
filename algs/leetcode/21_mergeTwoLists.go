package leetcode

// 23题的简化版，
// 思路1：23题的小顶堆解法可以直接挪过来
// 思路2，直接合并

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}

	cur, cur1, cur2 := head, l1, l2
	for ; cur1 != nil && cur2 != nil; cur = cur.Next {
		if cur1.Val > cur2.Val {
			cur.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		} else {
			cur.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		}
	}

	for ; cur1 != nil; cur = cur.Next {
		cur.Next = &ListNode{Val: cur1.Val}
		cur1 = cur1.Next
	}

	for ; cur2 != nil; cur = cur.Next {
		cur.Next = &ListNode{Val: cur2.Val}
		cur2 = cur2.Next
	}

	return head.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}

	cur, cur1, cur2 := head, l1, l2
	for ; cur1 != nil && cur2 != nil; cur = cur.Next {
		if cur1.Val > cur2.Val {
			cur.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		} else {
			cur.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		}
	}

	if cur1 != nil {
		cur.Next = cur1
	}

	if cur2 != nil {
		cur.Next = cur2
	}

	return head.Next
}
