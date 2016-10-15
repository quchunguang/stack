package stack

import "container/list"

type Stack struct {
	sem  chan int
	list *list.List
}

// NewStack create a new stack.
func New() *Stack {
	sem := make(chan int, 1)
	list := list.New()
	return &Stack{sem, list}
}

// Push item to stack.
func (s *Stack) Push(value interface{}) {
	s.sem <- 1
	s.list.PushBack(value)
	<-s.sem
}

// Pop item from stack.
func (s *Stack) Pop() interface{} {
	s.sem <- 1
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
	}
	<-s.sem

	if e != nil {
		return e.Value
	} else {
		return nil
	}
}

// Peak get the top item of the stack.
func (s *Stack) Peak() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len get the length of the stack.
func (s *Stack) Len() int {
	return s.list.Len()
}

// Empty tests if the stack is empty.
func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
