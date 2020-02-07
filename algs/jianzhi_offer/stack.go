package jianzhi_offer

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return s[len(s)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	oldStack := *s
	if len(oldStack) == 0 {
		return nil, errors.New("out of index, len is 0")
	}

	ret := oldStack[len(oldStack)-1]
	*s = oldStack[:len(oldStack)-1]
	return ret, nil
}
