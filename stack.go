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
func (stack *Stack) Push(value interface{}) {
	stack.sem <- 1
	stack.list.PushBack(value)
	<-stack.sem
}

// Pop item from stack.
func (stack *Stack) Pop() interface{} {
	stack.sem <- 1
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
	}
	<-stack.sem

	if e != nil {
		return e.Value
	} else {
		return nil
	}
}

// Peak get the top item of the stack.
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len get the length of the stack.
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty tests if the stack is empty.
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
