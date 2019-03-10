package tree

import "github.com/adrianosela/GoTools/primitives/ints"

// isBalanced returns whether a node is balanced
func isBalanced(n *AVLNode) bool {
	// a leaf is balanced
	if n == nil || (n.Left == nil && n.Right == nil) {
		return true
	}
	// if one of the two is nil, other must have Height
	// less than or equal to 1 to be balanced
	if n.Left != nil && n.Right == nil {
		return n.Left.Height <= 1
	}
	if n.Left == nil && n.Right != nil {
		return n.Right.Height <= 1
	}
	// otherwise the absolute value of the difference in subtree heights
	// must be at most 1
	return ints.Abs(n.Left.Height-n.Right.Height) <= 1
}

// updateHeight recursively computes and returns the height at node n
func updateHeight(n *AVLNode) int {
	if n == nil || (n.Right == nil && n.Left == nil) {
		return 0
	}
	n.Height = 1 + ints.Max(updateHeight(n.Left), updateHeight(n.Right))
	return n.Height
}

// rotateLeft performs a left roation on a node
func rotateLeft(a *AVLNode) {
	b := a.Right
	a.Right = b.Left
	b.Left = a
	updateHeight(a)
	updateHeight(b)
	*a = *b // fixme this might not be correct
}

// rotateLeft performs a right roation on a node
func rotateRight(b *AVLNode) {
	a := b.Left
	b.Left = a.Right
	a.Right = b
	updateHeight(b)
	updateHeight(a)
	*b = *a // fixme this might not be correct
}

// rotateLeftRight performs a left roation on a node's left node, and then
// a right notation on the node itself
func rotateLeftRight(n *AVLNode) {
	rotateLeft(n.Left)
	rotateRight(n)
}

// rotateRightLeft performs a right roation on a node's right node, and then
// a left notation on the node itself
func rotateRightLeft(n *AVLNode) {
	rotateRight(n.Right)
	rotateLeft(n)
}
