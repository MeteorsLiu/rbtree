package rbtree

import "cmp"

// -1 if x is less than y,
//
//	0 if x equals y,
//
// +1 if x is greater than y.
type Comparator[K cmp.Ordered] func(x, y K) int

type InsertFn[K cmp.Ordered, V any] func(*Tree[K, V], K, V) *Node[K, V]

func WithDefaultComparator[K cmp.Ordered, V any]() Options[K, V] {
	return func(tree *Tree[K, V]) {
		tree.Comparator = cmp.Compare
	}
}

func defaultInsert[K cmp.Ordered, V any](tree *Tree[K, V], key K, value V) (insertedNode *Node[K, V]) {
	node := tree.Root
	for {
		ret := tree.Comparator(key, node.Key)
		// a > b
		if ret > 0 {
			if node.Right == nil {
				node.Right = NewNode(key, value)
				insertedNode = node.Right
				break
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = NewNode(key, value)
				insertedNode = node.Left
				break
			}
			node = node.Left
		}
	}
	insertedNode.SetParent(node)
	return
}

func WithDefaultInsert[K cmp.Ordered, V any]() Options[K, V] {
	return func(tree *Tree[K, V]) {
		tree.InsertFn = defaultInsert
	}
}
