package rbtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BST[K constraints.Ordered, T any] struct {
	root *RBNode[K, T]
}

func NewBSTree[K constraints.Ordered, T any]() *BST[K, T] {
	var zero K
	nilNode := NewRBNode[K, T](nil, nil, nil, zero)
	return &BST[K, T]{nilNode}
}

func (t *BST[K, T]) Insert(key K, data T) {
	node := t.root
	for {
		if key < node.key {
			if node.left == nil {
				node.left = NewRBNode[K, T](node, nil, nil, key, data)
				break
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = NewRBNode[K, T](node, nil, nil, key, data)
				break
			}
			node = node.right
		}
	}
}

func (t *BST[K, T]) print(n *RBNode[K, T], space int) {
	if n == nil {
		return
	}
	space += 10
	t.print(n.right, space)

	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%v => %v\n", n.Key(), n.Data())

	t.print(n.left, space)
}

func (t *BST[K, T]) Print() {
	t.print(t.root, 0)
}

func (t *BST[K, T]) Search(key K) *RBNode[K, T] {
	node := t.root

	for node != nil && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node
}

func (t *BST[K, T]) Min(n *RBNode[K, T]) *RBNode[K, T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *BST[K, T]) Leftmost(n *RBNode[K, T]) *RBNode[K, T] {
	return t.Min(t.root)
}
