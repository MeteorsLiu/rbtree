package rbtree

import "cmp"

type Color uint8

const (
	RED Color = iota
	BLACK
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

func (n *Node[K, V]) Color() Color {
	if n == nil {
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

func (c Color) String() string {
	if c == RED {
		return "RED"
	}
	return "BLACK"
}
