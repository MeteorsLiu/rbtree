package rbtree

import (
	"cmp"
	"fmt"
)

// Node is a single element within the tree
type Node[K cmp.Ordered, V any] struct {
	Key        K
	Value      V
	color      Color
	Left       *Node[K, V]
	Right      *Node[K, V]
	parent     *Node[K, V]
	isSentinel bool
}

func NewSentinel[K cmp.Ordered, V any]() *Node[K, V] {
	return &Node[K, V]{color: BLACK, isSentinel: true}
}

func NewNode[K cmp.Ordered, V any](sentinel *Node[K, V], key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key:    key,
		Value:  value,
		Left:   sentinel,
		Right:  sentinel,
		parent: sentinel,
	}
}

// RemoveFrom removes a node from a tree.
// it requires node to be removed MUST be in the tree.
func (node *Node[K, V]) RemoveFrom(tree *Tree[K, V]) {
	// no need to remove
	if node == nil || tree.Empty() {
		return
	}
	var temp, subst *Node[K, V]

	if node.Left.IsNil() {
		temp = node.Right
		subst = node
	} else if node.Right.IsNil() {
		temp = node.Left
		subst = node
	} else {
		subst = node.Right.min()
		temp = subst.Right
	}
	defer tree.decr()

	if subst == tree.Root {
		tree.Root = temp
		tree.Root.SetColor(BLACK)
		return
	}

	isBlack := subst.IsBlack()

	if subst == subst.Parent().LeftChild() {
		subst.Parent().SetLeftChild(temp)
	} else {
		subst.Parent().SetRightChild(temp)
	}

	if subst != node {
		if subst.Parent() == node {
			temp.SetParent(subst)
		} else {
			temp.SetParent(subst.Parent())
		}

		subst.SetLeftChild(node.LeftChild())
		subst.SetRightChild(node.RightChild())
		subst.SetParent(node.Parent())
		subst.CopyColorFrom(node)

		if node == tree.Root {
			tree.Root = subst
		} else {
			if node == node.Parent().LeftChild() {
				node.Parent().SetLeftChild(subst)
			} else {
				node.Parent().SetRightChild(subst)
			}
		}

		if !subst.LeftChild().IsNil() {
			subst.Left.SetParent(subst)
		}
		if !subst.RightChild().IsNil() {
			subst.Right.SetParent(subst)
		}
	} else {
		temp.SetParent(subst.Parent())
	}

	if isBlack {
		tree.fixup(temp)
	}
}

// IsNil indicates a node is a "nil" node or not.
// "nil" can be sentinel node or go's nil pointer.
func (n *Node[K, V]) IsNil() bool {
	return n == nil || n.isSentinel
}

func (n *Node[K, V]) SetColor(color Color) {
	n.color = color
}

func (n *Node[K, V]) SetParent(node *Node[K, V]) {
	n.parent = node
}

func (n *Node[K, V]) SetLeftChild(node *Node[K, V]) {
	n.Left = node
}
func (n *Node[K, V]) SetRightChild(node *Node[K, V]) {
	n.Right = node
}

// Copy node's color to n.
func (n *Node[K, V]) CopyColorFrom(node *Node[K, V]) {
	n.SetColor(node.Color())
}

// Color returns the color of the node
// color of "nil" nodes are always black.
func (n *Node[K, V]) Color() Color {
	// Sentinel may be changed sometime,
	// so we need to make sure its color always black
	if n.IsNil() {
		return BLACK
	}
	return n.color
}

func (n *Node[K, V]) IsRed() bool {
	return n.Color() == RED
}

func (n *Node[K, V]) IsBlack() bool {
	return n.Color() == BLACK
}

func (n *Node[K, V]) LeftChild() *Node[K, V] {
	return n.Left
}

func (n *Node[K, V]) RightChild() *Node[K, V] {
	return n.Right
}

func (n *Node[K, V]) Parent() *Node[K, V] {
	return n.parent
}

func (n *Node[K, V]) Grandparent() *Node[K, V] {
	return n.Parent().Parent()
}

func (n *Node[K, V]) min() *Node[K, V] {
	node := n
	for !node.Left.IsNil() {
		node = node.Left
	}
	return node
}
func (n *Node[K, V]) max() *Node[K, V] {
	node := n
	for !node.Right.IsNil() {
		node = node.Right
	}
	return node
}

// Min returns the minimum(leftmost) node of the node.
//
// If the result is "nil" node, it will return Go's nil pointer instead.
func (n *Node[K, V]) Min() *Node[K, V] {
	node := n.min()
	if node.IsNil() {
		return nil
	}
	return node
}

// Max returns the maximum(rightmost) node of the node.
//
// If the result is "nil" node, it will return Go's nil pointer instead.
func (n *Node[K, V]) Max() *Node[K, V] {
	node := n.max()
	if node.IsNil() {
		return nil
	}
	return node
}

func (node *Node[K, V]) String() string {
	return fmt.Sprintf("%v: %v: %s", node.Key, node.Value, node.Color())
}
