package stacks

import "errors"

type StackOfInt struct {
	values []int
}

func (s *StackOfInt) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInt) Push(x int) {
	s.values = append(s.values, x)
}

func (s *StackOfInt) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	idx := len(s.values) - 1
	val := s.values[idx]
	s.values = s.values[:idx]
	return val, nil

}
