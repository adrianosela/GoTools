package stack

// ArrStack is an array implementation of a stack
type ArrStack struct {
	c    []interface{}
	top  int
	size int
}

// NewArrStack is the constructor for an array impl. of a stack
func NewArrStack() *ArrStack {
	return &ArrStack{
		top:  -1,
		c:    make([]interface{}, 1),
		size: 1,
	}
}

// Push adds an item to the top of the stack
// Constant runtime O(1) UNTIL FULL
// Linear runtime O(n) averaged over n pushes starting from a size 0 stack
func (s *ArrStack) Push(data interface{}) {
	if s.isFull() {
		// built-in append func automatically resizes
		// but dont be fooled, that takes time
		s.c = append(s.c, data)
		s.top++
		s.size++
	} else {
		s.top++
		s.c[s.top] = data
	}
}

// Peek returns the item at the top of the stack but does not remove it
// Constant runtime O(1)
func (s *ArrStack) Peek() interface{} {
	return s.c[s.top]
}

// Pop returns the item at the top of the stack and returns it
// Constant runtime O(1)
func (s *ArrStack) Pop() interface{} {
	p := s.c[s.top]
	// we decrement top but dont downsize
	s.top--
	return p
}

func (s *ArrStack) isFull() bool {
	return (s.top == s.size)
}
