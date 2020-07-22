package algorithms

import "fmt"

const overflowLabel = -1

type stack struct {
	top   int
	stack []uint32
	size  uint32
}

func NewStack(size uint32) *stack {
	internalStack := make([]uint32, size)
	return &stack{
		top:   int(size),
		stack: internalStack,
		size:  size,
	}
}

func (s *stack) Push(element uint32) error {
	if s.top == overflowLabel {
		return fmt.Errorf("stack overflow")
	}
	s.top--
	s.stack[s.top] = element
	return nil
}

func (s *stack) Pop() (uint32, error) {
	if s.IsEmpty() {
		errMsg := fmt.Errorf("stack is empty")
		return 0, errMsg
	}
	result := s.stack[s.top]
	s.top++
	return result, nil
}

func (s *stack) IsEmpty() bool {
	return uint32(s.top) == s.size
}
