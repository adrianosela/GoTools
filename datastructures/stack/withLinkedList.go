package stack

import "github.com/adrianosela/GoTools/datastructures/linkedlist"

// LLStack is a linked-list implementation of a stack
type LLStack struct {
	top *linkedlist.Node
}

// NewLLStack is the constructor for the
// linked-list implementation of a stack
func NewLLStack() *LLStack {
	return &LLStack{}
}

// Push adds an item to the top of the stack
// Constant runtime O(1)
func (s *LLStack) Push(data interface{}) {
	s.top = &linkedlist.Node{
		Data: data,
		Next: s.top,
	}
}

// Peek returns the item at the top of the stack but does not remove it
// Constant runtime O(1)
func (s *LLStack) Peek() interface{} {
	if !s.isEmpty() {
		return s.top.Data
	}
	return nil
}

// Pop returns the item at the top of the stack and returns it
// Constant runtime O(1)
func (s *LLStack) Pop() interface{} {
	popped := s.top.Data
	s.top = s.top.Next
	return popped
}

func (s *LLStack) isEmpty() bool {
	return s.top == nil
}
