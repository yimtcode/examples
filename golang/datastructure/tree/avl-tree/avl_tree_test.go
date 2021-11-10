package avl_tree

import "testing"

func TestNewAVLTree(t *testing.T) {
	tree := new(AVLTree)
	//tree.Add(1)
	//tree.Add(2)
	tree.Add(3)

	d := tree.Find(3)
	t.Log(*d)

	tree.Remove(3)
}
