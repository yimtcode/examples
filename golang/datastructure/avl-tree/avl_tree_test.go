package avl_tree

import "testing"

func TestNewAVLTree(t *testing.T) {
	compare := func(d1, d2 interface{}) int {
		i1 := d1.(int)
		i2 := d2.(int)
		if i1 < i2 {
			return -1
		} else if i1 == i2 {
			return 0
		} else {
			return 1
		}
	}
	tree := new(AVLTree).SetCompare(compare)
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
}
