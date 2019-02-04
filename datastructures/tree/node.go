package tree

// Node is a generic "node" e.g. for a binary tree
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}
