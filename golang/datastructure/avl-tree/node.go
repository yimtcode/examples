package avl_tree

import "math"

type Node struct {
	Data        interface{}
	Left, Right *Node
	height      int
}

func (n *Node) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) isBalanced() bool {
	if n == nil {
		return true
	}

	height := int(math.Abs(float64(n.Left.GetHeight() - n.Right.GetHeight())))
	if height > 1 {
		return false
	}

	return n.Left.isBalanced() && n.Right.isBalanced()
}

func (n *Node)rightRotate() {
}