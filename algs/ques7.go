package algs

import "errors"

// 两个栈实现一个队列
type MyQueue struct {
	stack1, stack2 *Stack
}

func (q *MyQueue) IsEmpty() bool {
	return q.size() == 0
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		stack1: &Stack{},
		stack2: &Stack{},
	}
}

func (q *MyQueue) enqueue(val interface{}) {
	q.stack1.Push(val)
}

func (q *MyQueue) dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			val, _ := q.stack1.Pop()
			q.stack2.Push(val)
		}
	}

	ret, _ := q.stack2.Pop()
	return ret, nil
}

func (q *MyQueue) size() int {
	return q.stack1.Len() + q.stack2.Len()
}

type Queue interface {
	enqueue(val interface{})
	dequeue() (interface{}, error)
	size() int
	IsEmpty() bool
}

// 两个队列实现一个栈
type MyStack struct {
	queue1, queue2 *MyQueue
}

func NewMyStack() *MyStack {
	return &MyStack{
		queue1: NewMyQueue(),
		queue2: NewMyQueue(),
	}
}

func (stack *MyStack) IsEmpty() bool {
	return stack.size() == 0
}

func (stack *MyStack) Push(val interface{}) {
	if stack.queue1.IsEmpty() {
		stack.queue2.enqueue(val)
	} else {
		stack.queue1.enqueue(val)
	}
}

func (stack *MyStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	if stack.queue1.IsEmpty() {
		for stack.queue2.size() != 1 {
			val, _ := stack.queue2.dequeue()
			stack.queue1.enqueue(val)
		}
		return stack.queue2.dequeue()
	} else {
		for stack.queue1.size() != 1 {
			val, _ := stack.queue1.dequeue()
			stack.queue2.enqueue(val)
		}
		return stack.queue1.dequeue()
	}
}

func (stack *MyStack) size() int {
	return stack.queue1.size() + stack.queue2.size()
}

type IStack interface {
	Push(val interface{})
	Pop() (interface{}, error)
	size() int
	IsEmpty() bool
}
