package jianzhi_offer

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return
}

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/
// 找差值，找到交点前共同长度的起点，一起遍历找交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cura, curb := headA, headB

	var lena, lenb int
	for ; cura != nil; lena++ {
		cura = cura.Next
	}

	for ; curb != nil; lenb++ {
		curb = curb.Next
	}

	if lena > lenb {
		dif := lena - lenb
		for cura = headA; dif > 0; dif-- {
			cura = cura.Next
		}
	} else {
		dif := lenb - lena
		for curb = headB; dif > 0; dif-- {
			curb = curb.Next
		}
	}

	for ; cura != nil && curb != nil; cura, curb = cura.Next, curb.Next {
		if cura == curb {
			return cura
		}
	}

	return nil
}
