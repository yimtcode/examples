package avl_tree

func rightRotate(node *Node) *Node {
	newNode := node.Left
	newNode.Right = node
	return newNode
}
