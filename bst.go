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
