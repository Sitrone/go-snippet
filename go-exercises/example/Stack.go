package example

import (
	"errors"
	"fmt"
)

type stack interface {
	Push(i int)
	Pop() int
	IsEmpty() bool
}

type Stack struct {
	n   int
	arr []int
}

func NewStack() *Stack {
	return &Stack{
		n:   0,
		arr: make([]int, 0),
	}
}

func (s *Stack) Push(i int) {
	s.arr = append(s.arr, i)
	s.n++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	ret := s.arr[s.n-1]
	s.arr = s.arr[:s.n-1]
	s.n--
	return ret, nil
}

func (s *Stack) IsEmpty() bool {
	return s.n == 0
}

func (s *Stack) resize() {
	t := make([]int, len(s.arr), (cap(s.arr)+1)*2) // +1 in case cap(s) == 0
	for i := range s.arr {
		t[i] = s.arr[i]
	}
	s.arr = t
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.arr)
}
