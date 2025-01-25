package rbtree

import "cmp"

// LookupFn is the custom look up function for the redblack tree.
// the tree will use this function to Get() and Remove()
// if the node is not found, it's required to return Go's nil pointer to indicate it.
//
// Example please read the internal function "defaultLookup" in the current file.
type LookupFn[K cmp.Ordered, V any] func(*Node[K, V], K) (node *Node[K, V])

// InsertLookupFn is the custom look up function for tree inserting.
// the tree will use this function to Insert().
// it's required to indicate the place where tree can insert the node into indirectly.
//
// Example please read the internal function "defaultInsertLookup" in the current file.
type InsertLookupFn[K cmp.Ordered, V any] func(*Node[K, V], K) (indirect **Node[K, V], parent *Node[K, V])

func defaultInsertLookup[K cmp.Ordered, V any](root *Node[K, V], key K) (indirect **Node[K, V], parent *Node[K, V]) {
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

func defaultLookup[K cmp.Ordered, V any](root *Node[K, V], key K) (result *Node[K, V]) {
	node := root
	for !node.IsNil() {
		switch cmp.Compare(key, node.Key) {
		case -1:
			// a < b
			node = node.Left
		case +1:
			// a > b
			node = node.Right
		default:
			// a == b, found
			result = node
			return
		}
	}
	return
}

func WithDefaultInsertLookupFn[K cmp.Ordered, V any]() Options[K, V] {
	return func(tree *Tree[K, V]) {
		tree.InsertLookupFn = defaultInsertLookup
	}
}

func WithDefaultLookupFn[K cmp.Ordered, V any]() Options[K, V] {
	return func(tree *Tree[K, V]) {
		tree.LookupFn = defaultLookup
	}
}
