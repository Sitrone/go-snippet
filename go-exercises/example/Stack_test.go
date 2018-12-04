package example

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	stack.Push(3)
	fmt.Println(stack)
}