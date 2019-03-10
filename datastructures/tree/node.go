package tree

// Node is a generic "node" e.g. for a binary tree
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

// AVLNode is a node for an AVL tree implementation
type AVLNode struct {
	Data   interface{}
	Left   *AVLNode
	Right  *AVLNode
	Height int
}
