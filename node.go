package rbtree

import (
	"cmp"
)

// Node is a single element within the tree
type Node[K cmp.Ordered, V any] struct {
	Key    K
	Value  V
	color  Color
	Left   *Node[K, V]
	Right  *Node[K, V]
	parent *Node[K, V]
}

func NewNode[K cmp.Ordered, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{Key: key, Value: value}
}

func (node *Node[K, V]) RemoveFrom(tree *Tree[K, V]) {
	var child *Node[K, V]

	if node.Left != nil && node.Right != nil {
		pred := node.Left.Max()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.IsBlack() {
			node.CopyColorFrom(child)
			tree.fixup(node)
		}
		tree.replaceNode(node, child)
		if node.Parent() == nil && child != nil {
			child.SetColor(BLACK)
		}
	}
	tree.size--
}

func (n *Node[K, V]) SetColor(color Color) {
	if n == nil {
		panic("warn: set a nil node")
	}
	n.color = color
}

func (n *Node[K, V]) SetParent(node *Node[K, V]) {
	if n == nil {
		panic("warn: set a nil node")
	}
	n.parent = node
}

func (n *Node[K, V]) SetLeftChild(node *Node[K, V]) {
	if n == nil {
		panic("warn: set a nil node")
	}
	n.Left = node
}
func (n *Node[K, V]) SetRightChild(node *Node[K, V]) {
	if n == nil {
		panic("warn: set a nil node")
	}
	n.Right = node
}

func (n *Node[K, V]) Color() Color {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (n *Node[K, V]) CopyColorFrom(node *Node[K, V]) {
	n.SetColor(node.Color())
}

func (n *Node[K, V]) IsRed() bool {
	return n.Color() == RED
}

func (n *Node[K, V]) IsBlack() bool {
	return n.Color() == BLACK
}

func (n *Node[K, V]) LeftChild() *Node[K, V] {
	if n == nil {
		return nil
	}
	return n.Left
}

func (n *Node[K, V]) RightChild() *Node[K, V] {
	if n == nil {
		return nil
	}
	return n.Right
}

func (n *Node[K, V]) Parent() *Node[K, V] {
	if n == nil {
		return nil
	}
	return n.parent
}

func (n *Node[K, V]) Grandparent() *Node[K, V] {
	if n == nil {
		return nil
	}
	return n.Parent().Parent()
}

func (n *Node[K, V]) Min() *Node[K, V] {
	if n == nil {
		return nil
	}
	node := n
	for node.Left != nil {
		node = node.Left
	}
	return node
}
func (n *Node[K, V]) Max() *Node[K, V] {
	if n == nil {
		return nil
	}
	node := n
	for node.Right != nil {
		node = node.Right
	}
	return node
}
