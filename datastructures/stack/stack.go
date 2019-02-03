package stack

// Stack represents the generic stack interface
type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
