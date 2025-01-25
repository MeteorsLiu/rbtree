package rbtree

import "cmp"

type LookupFn[K cmp.Ordered, V any] func(*Node[K, V], K) (indirect **Node[K, V], parent *Node[K, V])

func defaultLookup[K cmp.Ordered, V any](root *Node[K, V], key K) (indirect **Node[K, V], parent *Node[K, V]) {
	node := root
	for {
		// a < b
		if cmp.Less(key, node.Key) {
			if node.Left.IsNil() {
				indirect = &node.Left
				break
			}
			node = node.Left
		} else {
			if node.Right.IsNil() {
				indirect = &node.Right
				break
			}
			node = node.Right
		}
	}
	parent = node
	return
}

func WithDefaultLookupFn[K cmp.Ordered, V any]() Options[K, V] {
	return func(tree *Tree[K, V]) {
		tree.LookupFn = defaultLookup
	}
}
