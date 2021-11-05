package avl_tree

type AVLTree struct {
	root Node
	size int
}

func NewAVLTree(data interface{}) *AVLTree {
	return &AVLTree{
		root: Node{Data: data, height: 1},
		size: 1,
	}
}

func (a *AVLTree) GetSize() int {
	return a.size
}

func (a *AVLTree) IsEmpty() bool {
	return a.size == 0
}

func (a *AVLTree) isBalanced() bool {
	return a.root.isBalanced()
}
