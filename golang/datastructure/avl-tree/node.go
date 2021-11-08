package avl_tree

import "math"

type Node struct {
	Data        interface{}
	left, right *Node
	height      int
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) isBalanced() bool {
	if n == nil {
		return true
	}

	height := int(math.Abs(float64(n.left.getHeight() - n.right.getHeight())))
	if height > 1 {
		return false
	}

	return n.left.isBalanced() && n.right.isBalanced()
}

func (n *Node) rightRotate() *Node {
	top := n.left
	n.left = top.right
	top.right = n
	// update height
	top.height = int(math.Max(float64(top.left.getHeight()), float64(top.right.getHeight())))
	n.height = int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight())))
	return top
}

func (n *Node) leftRotate() *Node {
	top := n.right
	n.right = top.left
	top.left = n
	// update height
	top.height = int(math.Max(float64(top.left.getHeight()), float64(top.right.getHeight())))
	n.height = int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight())))
	return top
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data, height: 1}
}

func (n *Node) getBalanceFactor() int {
	return n.left.getHeight() - n.right.getHeight()
}

func (n *Node) add(data interface{}, compare func(d1, d2 interface{}) int) {
	r := compare(data, n.Data)
	if r <= 0 {
		// less and equal
		if n.left == nil {
			n.left = NewNode(data)
		} else {
			n.left.add(data, compare)
		}
	} else if r > 0 {
		// greater
		if n.right == nil {
			n.right = NewNode(data)
		} else {
			n.right.add(data, compare)
		}
	}

	n.height = 1 + int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight()))
	factor := n.getBalanceFactor()
}
