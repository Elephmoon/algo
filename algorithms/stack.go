package algorithms

import "fmt"

type stack struct {
	top   int
	stack []uint32
}

func NewStack() *stack {
	return &stack{
		top:   0,
		stack: []uint32{},
	}
}

func (s *stack) Push(element uint32) {
	if len(s.stack) != 0 {
		s.top++
	}
	s.stack = append(s.stack, element)
}

func (s *stack) Pop() (uint32, error) {
	if s.IsEmpty() {
		errMsg := fmt.Errorf("stack is empty")
		return 0, errMsg
	}
	result := s.stack[s.top]
	s.top--

	if s.top > 0 {
		_ = copy(s.stack, s.stack[:s.top])
	} else {
		s.stack = []uint32{}
	}

	return result, nil
}

func (s *stack) IsEmpty() bool {
	return len(s.stack) == 0
}
