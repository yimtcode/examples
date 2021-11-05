package avl_tree

import "testing"

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree(1)
	t.Log(tree.isBalanced())
}
