package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum, canary int
		ans         = &ListNode{Val: -1}
	)

	cur, cur1, cur2 := ans, l1, l2
	for ; cur1 != nil && cur2 != nil; cur, cur1, cur2 = cur.Next, cur1.Next, cur2.Next {
		sum = cur1.Val + cur2.Val + canary
		sum, canary = sum%10, sum/10

		cur.Next = &ListNode{Val: sum}
	}

	for ; cur1 != nil; cur, cur1 = cur.Next, cur1.Next {
		sum = cur1.Val + canary
		sum, canary = sum%10, sum/10
		cur.Next = &ListNode{Val: sum}
	}

	for ; cur2 != nil; cur, cur2 = cur.Next, cur2.Next {
		sum = cur2.Val + canary
		sum, canary = sum%10, sum/10
		cur.Next = &ListNode{Val: sum}
	}

	if canary > 0 {
		cur.Next = &ListNode{Val: canary}
	}

	return ans.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum, canary int
		ans         = &ListNode{Val: -1}
	)

	cur, cur1, cur2 := ans, l1, l2
	for ; cur1 != nil || cur2 != nil; cur = cur.Next {
		sum = canary
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}

		sum, canary = sum%10, sum/10
		cur.Next = &ListNode{Val: sum}
	}

	if canary > 0 {
		cur.Next = &ListNode{Val: canary}
	}

	return ans.Next
}
