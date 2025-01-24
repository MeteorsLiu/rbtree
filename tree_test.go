package rbtree

import "testing"

func TestRBTree(t *testing.T) {
	tree := NewTree[int, int]()

	for i := 0; i < 10; i++ {
		tree.Insert(i, i)
	}

	t.Log(tree.String())
}
