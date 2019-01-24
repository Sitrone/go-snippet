package algs

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	oldStack := *stack
	if len(oldStack) == 0 {
		return nil, errors.New("out of index, len is 0")
	}

	ret := oldStack[len(oldStack)-1]
	*stack = oldStack[:len(oldStack)-1]
	return ret, nil
}
