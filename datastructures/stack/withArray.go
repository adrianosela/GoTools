package stack

// ArrStack is an array implementation of a stack
type ArrStack struct {
	c   []interface{}
	top int
}

// Push adds an item to the top of the stack
func (s *ArrStack) Push(data interface{}) {
	if s.isFull() {
		// built-in append func automatically resizes
		// but dont be fooled, that takes time
		s.c = append(s.c, data)
		s.top++
	} else {
		s.top++
		s.c[s.top] = data
	}
}

// Peek returns the item at the top of the stack but does not remove it
func (s *ArrStack) Peek() interface{} {
	return s.c[s.top]
}

// Pop returns the item at the top of the stack and returns it
func (s *ArrStack) Pop() interface{} {
	p := s.c[s.top]
	// we decrement top but dont downsize
	s.top--
	return p
}

func (s *ArrStack) isFull() bool {
	return (s.top == len(s.c))
}
