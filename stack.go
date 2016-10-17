package stack

import "container/list"

// Stack is describe as a list with lock
type Stack struct {
	sem  chan int
	list *list.List
}

// CallbackFunc type describe any operation on each element in stack.
type CallbackFunc func(val interface{}) bool

// New create a new stack.
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

// Map returns the first element in the stack causing mapFunc returns true.
func (s *Stack) Map(mapFunc CallbackFunc) interface{} {
	s.sem <- 1
	e := s.list.Front()
	for e != nil {
		if mapFunc(e.Value) {
			<-s.sem
			return e.Value
		}
		e = e.Next()
	}
	<-s.sem
	return nil
}

// Contain tests if this item in the stack.
func (s *Stack) Contain(val interface{}) bool {
	s.sem <- 1
	e := s.list.Front()
	for e != nil {
		if e.Value == val {
			<-s.sem
			return true
		}
		e = e.Next()
	}
	<-s.sem
	return false
}
