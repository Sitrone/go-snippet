package algs

import (
	"fmt"
	"testing"
)

func TestMyQueue(t *testing.T) {
	var q = NewMyQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	q.enqueue(4)
	fmt.Println(q.dequeue())
}

func TestMyStack_Pop(t *testing.T) {
	var stack = NewMyStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push(4)
	stack.Push(5)
	fmt.Println(stack.Pop())
}