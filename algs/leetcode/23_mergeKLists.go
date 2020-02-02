package leetcode

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/merge-k-sorted-lists/

type ListNodeHeap []*ListNode

func (lp *ListNodeHeap) Push(x interface{}) {
	*lp = append(*lp, x.(*ListNode))
}

func (lp *ListNodeHeap) Pop() interface{} {
	old := *lp
	n := len(old)
	x := old[n-1]
	*lp = old[0 : n-1]
	return x
}

func (lp ListNodeHeap) Len() int {
	return len(lp)
}

func (lp ListNodeHeap) Less(i, j int) bool {
	return lp[i].Val < lp[j].Val
}

func (lp ListNodeHeap) Swap(i, j int) {
	lp[i], lp[j] = lp[j], lp[i]
}

// 思路1：小顶堆, 多路归并，对于大数据量的场景比较友好，时间复杂度O(nlogn)，空间复杂度O(m)
// 思路2：对于小数据量，可以直接全部取出，然后快排
// 思路3：两两合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	lp := &ListNodeHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(lp, lists[i])
		}
	}

	if lp.Len() == 0 {
		return nil
	}

	head := heap.Pop(lp).(*ListNode)
	cur := head
	nextToPush := cur.Next

	var next *ListNode
	for lp.Len() != 0 {
		if nextToPush != nil {
			heap.Push(lp, nextToPush)
		}
		next = heap.Pop(lp).(*ListNode)
		cur.Next = next
		cur = cur.Next
		nextToPush = next.Next
	}

	return head
}

// 排序
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	ret := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		var cur = lists[i]
		for cur != nil {
			ret = append(ret, cur.Val)
			cur = cur.Next
		}
	}

	if len(ret) == 0 {
		return nil
	}

	sort.Ints(ret)

	head := &ListNode{Val: ret[0]}
	var cur = head
	for i := 1; i < len(ret); i++ {
		cur.Next = &ListNode{Val: ret[i]}
		cur = cur.Next
	}

	return head
}
