package tree

// Tree represents a binary tree
type Tree struct {
	Root *Node
}

// PreOrderTraverse returns the elements of a tree in pre order traversal
func (t *Tree) PreOrderTraverse() []interface{} {
	var elements []interface{}
	var f func(*Node)

	f = func(r *Node) {
		if r == nil {
			return
		}
		elements = append(elements, r.Data)
		f(r.Left)
		f(r.Right)
	}

	f(t.Root)
	return elements
}

// InOrderTraverse returns the elements of a tree in in-order traversal
func (t *Tree) InOrderTraverse() []interface{} {
	var elements []interface{}
	var f func(*Node)

	f = func(r *Node) {
		if r == nil {
			return
		}
		f(r.Left)
		elements = append(elements, r.Data)
		f(r.Right)
	}

	f(t.Root)
	return elements
}

// PostOrderTraverse returns the elements of a tree in post order traversal
func (t *Tree) PostOrderTraverse() []interface{} {
	var elements []interface{}
	var f func(*Node)

	f = func(r *Node) {
		if r == nil {
			return
		}
		f(r.Left)
		f(r.Right)
		elements = append(elements, r.Data)
	}

	f(t.Root)
	return elements
}
