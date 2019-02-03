package stack

import "github.com/adrianosela/GoTools/datastructures/linkedlist"

// LLStack is a linked-list implementation of a stack
type LLStack struct {
	top *linkedlist.Node
}

// Push adds an item to the top of the stack
func (s *LLStack) Push(data interface{}) {
	s.top = &linkedlist.Node{
		Data: data,
		Next: s.top,
	}
}

// Peek returns the item at the top of the stack but does not remove it
func (s *LLStack) Peek() interface{} {
	return s.top.Data
}

// Pop returns the item at the top of the stack and returns it
func (s *LLStack) Pop() interface{} {
	popped := s.top.Data
	s.top = s.top.Next
	return popped
}
