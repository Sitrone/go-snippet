package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || k == 0 {
		return head
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	if k >= length {
		k %= length
	}

	if k == 0 {
		return head
	}

	k = length - k

	cur := head
	for ; cur != nil && k > 1; cur = cur.Next {
		k--
	}

	next := cur.Next
	newHead := next
	cur.Next = nil

	for next.Next != nil {
		next = next.Next
	}

	next.Next = head

	return newHead
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || k == 0 {
		return head
	}

	length := 1
	last := head
	for ; last.Next != nil; last = last.Next {
		length++
	}

	last.Next = head // 形成环

	if k >= length {
		k %= length
	}

	k = length - k

	newLast := head
	for ; newLast != nil && k > 1; newLast = newLast.Next {
		k--
	}

	newHead := newLast.Next
	newLast.Next = nil

	return newHead
}
