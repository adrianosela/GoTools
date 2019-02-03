package linkedlist

import "fmt"

// Node is a generic "node" e.g. for singly linked lists
type Node struct {
	Data interface{}
	Next *Node
}

// PNode adds a "previous" pointer to Node to allow doubly linked lists
type PNode struct {
	Node
	Prev *PNode
}

// NewNode is the constructor for the generic node object
func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

// WalkForward moves k steps along the next edge of nodes in a list
func (n *Node) WalkForward(k int) (*Node, error) {
	for i := 0; i < k; i++ {
		if n.Next != nil {
			n = n.Next
		} else {
			return nil, fmt.Errorf("list not long enough, found nil after %d steps", i)
		}
	}
	return n, nil
}

// WalkBackward moves k steps along the prev edge of nodes in a list
func (n *PNode) WalkBackward(k int) (*PNode, error) {
	for i := 0; i < k; i++ {
		if n.Prev != nil {
			n = n.Prev
		} else {
			return nil, fmt.Errorf("list not long enough, found nil after %d steps", i)
		}
	}
	return n, nil
}
