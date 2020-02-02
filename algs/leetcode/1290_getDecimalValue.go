package leetcode

//
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	num := head.Val
	for head.Next != nil {
		num = num << 1
		num |= head.Next.Val
		head = head.Next
	}

	return num
}
