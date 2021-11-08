package avl_tree

type AVLTree struct {
	root    *Node
	size    int
	compare func(d1, d2 interface{}) int
}

func (a *AVLTree) SetCompare(compare func(d1, d2 interface{}) int) *AVLTree {
	a.compare = compare
	return a
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

func (a *AVLTree) Add(data interface{}) {
	if a.root == nil {
		a.root = NewNode(data)
	} else {
		a.root.add(data, a.compare)
	}
}
