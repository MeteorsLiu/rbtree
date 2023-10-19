package rbtree

import (
	"cmp"
)

type Color uint8

const (
	RED Color = iota
	BLACK
)

func (c Color) String() string {
	if c == RED {
		return "RED"
	}
	return "BLACK"
}

const (
	LESS = iota - 1
	EQUAL
	GREATER
)

// this comparator sucks.
// However, we have to use it because when we need to use a custom comparator
// and this comparator is the fastest way to compare generics.
type Compare[K cmp.Ordered] func(a, b K) int

type RBNode[K cmp.Ordered, T any] struct {
	parent, left, right *RBNode[K, T]

	key  K
	data T

	color Color
}

func NewRBNode[K cmp.Ordered, T any](parent, left, right *RBNode[K, T], key K, data ...T) *RBNode[K, T] {
	if len(data) > 0 {
		return &RBNode[K, T]{
			key:    key,
			data:   data[0],
			parent: parent,
			left:   left,
			right:  right,
		}
	}
	return &RBNode[K, T]{
		key:    key,
		parent: parent,
		left:   left,
		right:  right,
	}
}

func (n *RBNode[K, T]) Reset(parent, left, right *RBNode[K, T], color Color) {
	n.parent = parent
	n.left = left
	n.right = right
	n.color = color
}

func (n *RBNode[K, T]) Color() Color {
	return n.color
}

func (n *RBNode[K, T]) Left() *RBNode[K, T] {
	return n.left
}

func (n *RBNode[K, T]) Right() *RBNode[K, T] {
	return n.right
}

func (n *RBNode[K, T]) Key() K {
	return n.key
}

func (n *RBNode[K, T]) Data() T {
	return n.data
}

func (n *RBNode[K, T]) Parent() *RBNode[K, T] {
	return n.parent
}

func (n *RBNode[K, T]) Grandparent() *RBNode[K, T] {
	return n.parent.parent
}

func (n *RBNode[K, T]) Uncle() *RBNode[K, T] {
	if n.Grandparent().left == n.parent {
		return n.Grandparent().right
	} else {
		return n.Grandparent().left
	}
}

func (n *RBNode[K, T]) Sibling() *RBNode[K, T] {
	if n.parent.left == n {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

func (n *RBNode[K, T]) IsBlack() bool {
	return n.color == BLACK
}

func (n *RBNode[K, T]) IsRed() bool {
	return n.color == RED
}

func (n *RBNode[K, T]) IsLeftChild() bool {
	return n.parent.left == n
}

func (n *RBNode[K, T]) IsRightChild() bool {
	return n.parent.right == n
}

func (n *RBNode[K, T]) SetColor(color Color) {
	n.color = color
}

//		 	  |                       |
//			  N                       S
//			 / \     l-rotate(N)     / \
//	   		L   S    ==========>    N   R
//		   	   / \                 / \
//		 	  M   R               L   M
func (t *RBTree[K, T]) RotateLeft(n *RBNode[K, T]) {
	right := n.right
	n.right = right.left

	if right.left != t.nilNode {
		right.left.parent = n
	}

	right.parent = n.parent

	switch {
	case n == t.root:
		t.root = right
	case n.IsLeftChild():
		n.parent.left = right
	default:
		n.parent.right = right
	}

	right.left = n
	n.parent = right
}

//			  |                       |
//			  N                       L
//			 / \     r-rotate(N)     / \
//		 	L   S    ==========>    M   N
//		   / \					       / \
//	   	  M   R						  R   S
func (t *RBTree[K, T]) RotateRight(n *RBNode[K, T]) {
	left := n.left
	n.left = left.right

	if left.right != t.nilNode {
		left.right.parent = n
	}

	left.parent = n.parent

	switch {
	case n == t.root:
		t.root = left
	case n.IsRightChild():
		n.parent.right = left
	default:
		n.parent.left = left
	}

	left.right = n
	n.parent = left
}
