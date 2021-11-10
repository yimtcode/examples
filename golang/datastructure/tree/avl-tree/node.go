package avl_tree

import "math"

type Node struct {
	Data        int
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
	// 注意先计算底部，在计算顶部
	n.height = int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight()))) + 1
	top.height = int(math.Max(float64(top.left.getHeight()), float64(top.right.getHeight()))) + 1
	return top
}

func (n *Node) leftRotate() *Node {
	top := n.right
	n.right = top.left
	top.left = n
	// 注意先计算底部，在计算顶部
	n.height = int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight()))) + 1
	top.height = int(math.Max(float64(top.left.getHeight()), float64(top.right.getHeight()))) + 1
	return top
}

func NewNode(data int) *Node {
	return &Node{Data: data, height: 1}
}

func (n *Node) getBalanceFactor() int {
	return n.left.getHeight() - n.right.getHeight()
}
