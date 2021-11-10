package avl_tree

import (
	"math"
)

type AVLTree struct {
	root *Node
	size int
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

func (a *AVLTree) Add(data int) {
	a.size++
	if a.root == nil {
		// 根节点没有数据
		a.root = NewNode(data)
		return
	}

	// 防止root发生旋转
	a.root = a.add(a.root, data)
}

func (a *AVLTree) Find(data int) *int {
	if a.IsEmpty() {
		// 空树
		return nil
	}

	n := a.find(a.root, data)
	if n == nil {
		return nil
	}
	t := n.Data
	return &t
}

func (a *AVLTree) Remove(data int) {
	if a.IsEmpty() {
		// 空树
		return
	}
	// 防止删除根节点
	a.root = a.remove(a.root, data)
}

// 内部使用
func (a *AVLTree) add(n *Node, data int) *Node {
	r := data - n.Data
	if r < 0 {
		// 左
		if n.left == nil {
			n.left = NewNode(data)
		} else {
			a.add(n.left, data)
		}
	} else if r > 0 {
		// 右
		if n.right == nil {
			n.right = NewNode(data)
		} else {
			a.add(n.right, data)
		}
	}

	// 平衡
	return a.balance(n)
}

func (a *AVLTree) balance(n *Node) *Node {
	// 更新高度
	n.height = 1 + int(math.Max(float64(n.left.getHeight()), float64(n.right.getHeight())))
	factor := n.getBalanceFactor()
	if factor > 1 && n.left.getBalanceFactor() > 0 {
		// LL
		return n.rightRotate()
	}

	if factor < -1 && n.right.getBalanceFactor() < 0 {
		// RR
		return n.leftRotate()
	}
	if factor > 1 && n.left.getBalanceFactor() < 0 {
		// LR
		t := n.left.rightRotate()
		return t.rightRotate()
	}
	if factor < -1 && n.right.getBalanceFactor() > 0 {
		// RL
		t := n.right.right
		return t.leftRotate()
	}

	return n
}

func (a *AVLTree) find(n *Node, data int) *Node {
	if n == nil {
		return nil
	}

	factor := data - n.Data
	if factor == 0 {
		return n
	}

	if factor < 0 {
		return a.find(n.left, data)
	} else {
		return a.find(n.right, data)
	}
}

// 返回被删除节点的替换节点
func (a *AVLTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil
	}
	var returnNode *Node
	factor := data - n.Data
	if factor < 0 {
		// 左
		n.left = a.remove(n.left, data)
		returnNode = n
	} else if factor > 0 {
		// 右
		n.right = a.remove(n.right, data)
		returnNode = n
	} else {
		// 找到需要删除
		a.size--
		if n.left == nil {
			// 左边为空
			returnNode = n.right
			n.right = nil
		} else if n.right == nil {
			// 删除右边为空
			returnNode = n.left
			n.left = nil
		} else {
			// 删除左右都不为空
			min := a.minimum(n.right)
			min.right = a.remove(n.right, min.Data)
			min.left = n.left
			n.right = nil
			n.left = nil
			returnNode = min
		}
	}
	if returnNode == nil {
		return nil
	}
	return a.balance(returnNode)
}

func (a *AVLTree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}

	return a.minimum(node.left)
}
