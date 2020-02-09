package leetcode

//https://leetcode-cn.com/problems/design-circular-deque/

// accepted 16ms
//type MyCircularDeque struct {
//	n     int
//	ele   []int
//	total int
//}
//
///** Initialize your data structure here. Set the size of the deque to be k. */
//func Constructor(k int) MyCircularDeque {
//	return MyCircularDeque{ele: make([]int, 0, k), total: k}
//}
//
///** Adds an item at the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertFront(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//
//	this.n++
//	this.ele = append([]int{value}, this.ele...)
//	return true
//}
//
///** Adds an item at the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertLast(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//
//	this.n++
//	this.ele = append(this.ele, value)
//	return true
//}
//
///** Deletes an item from the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteFront() bool {
//	if this.IsEmpty() {
//		return false
//	}
//
//	this.n--
//	this.ele = this.ele[1:]
//	return true
//}
//
///** Deletes an item from the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteLast() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	this.n--
//	this.ele = this.ele[:len(this.ele)-1]
//	return true
//}
//
///** Get the front item from the deque. */
//func (this *MyCircularDeque) GetFront() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.ele[0]
//}
//
///** Get the last item from the deque. */
//func (this *MyCircularDeque) GetRear() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.ele[len(this.ele)-1]
//}
//
///** Checks whether the circular deque is empty or not. */
//func (this *MyCircularDeque) IsEmpty() bool {
//	return this.n == 0
//}
//
///** Checks whether the circular deque is full or not. */
//func (this *MyCircularDeque) IsFull() bool {
//	return this.n == this.total
//}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

// accepted 20ms
type MyCircularDeque struct {
	len, cap  int
	dummyNode *ListNode2 // 哨兵节点
}

type ListNode2 struct {
	Val       int
	Next, Pre *ListNode2
}

/** Initialize your data structure here. Set the size of the deque to be k. */
//func Constructor(k int) MyCircularDeque {
//	node := &ListNode2{Val: -1}
//	node.Next = node
//	node.Pre = node
//	return MyCircularDeque{
//		len:       0,
//		cap:       k,
//		dummyNode: node,
//	}
//}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.len++
	node := &ListNode2{
		Val:  value,
		Next: this.dummyNode.Next,
		Pre:  this.dummyNode,
	}
	this.dummyNode.Next.Pre = node
	this.dummyNode.Next = node
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.len++
	node := &ListNode2{
		Val:  value,
		Next: this.dummyNode,
		Pre:  this.dummyNode.Pre,
	}
	this.dummyNode.Pre.Next = node
	this.dummyNode.Pre = node
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.len--
	this.dummyNode.Next.Next.Pre = this.dummyNode
	this.dummyNode.Next = this.dummyNode.Next.Next
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.len--
	this.dummyNode.Pre.Pre.Next = this.dummyNode
	this.dummyNode.Pre = this.dummyNode.Pre.Pre
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.dummyNode.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.dummyNode.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}
