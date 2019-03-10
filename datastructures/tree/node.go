package tree

// Node is a generic "node" e.g. for a binary tree
// height included for AVL tree implementations
type Node struct {
	Data  interface{}
	Height int
	Left  *Node
	Right *Node
}
