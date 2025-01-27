package rbtree

import (
	"cmp"
)

type Tree[K cmp.Ordered, V any] struct {
	size           int
	Root           *Node[K, V]
	sentinel       *Node[K, V]
	LookupFn       LookupFn[K, V]
	InsertLookupFn InsertLookupFn[K, V]
}

type Options[K cmp.Ordered, V any] func(*Tree[K, V])

func NewTree[K cmp.Ordered, V any](opts ...Options[K, V]) *Tree[K, V] {
	t := &Tree[K, V]{
		sentinel:       NewSentinel[K, V](),
		LookupFn:       defaultLookup[K, V],
		InsertLookupFn: defaultInsertLookup[K, V],
	}
	t.Root = t.sentinel

	for _, o := range opts {
		o(t)
	}

	return t
}

func (tree *Tree[K, V]) incr() {
	tree.size++
}

func (tree *Tree[K, V]) decr() {
	tree.size--
}

func (tree *Tree[K, V]) leftRotate(node *Node[K, V]) {
	right := node.RightChild()
	node.Right = right.LeftChild()

	if !right.LeftChild().IsNil() {
		right.Left.SetParent(node)
	}

	right.SetParent(node.Parent())

	if node == tree.Root {
		tree.Root = right
	} else if node == node.Parent().LeftChild() {
		node.Parent().SetLeftChild(right)
	} else {
		node.Parent().SetRightChild(right)
	}

	right.SetLeftChild(node)
	node.SetParent(right)
}

func (tree *Tree[K, V]) rightRotate(node *Node[K, V]) {
	left := node.LeftChild()
	node.Left = left.RightChild()

	if !left.RightChild().IsNil() {
		left.Right.SetParent(node)
	}

	left.SetParent(node.Parent())

	if node == tree.Root {
		tree.Root = left
	} else if node == node.Parent().RightChild() {
		node.Parent().SetRightChild(left)
	} else {
		node.Parent().SetLeftChild(left)
	}

	left.SetRightChild(node)
	node.SetParent(left)
}

func (tree *Tree[K, V]) fixup(node *Node[K, V]) {
	for node != tree.Root && node.IsBlack() {
		if node == node.Parent().LeftChild() {
			right := node.Parent().RightChild()

			if right.IsRed() {
				right.SetColor(BLACK)
				node.Parent().SetColor(RED)
				tree.leftRotate(node.Parent())
				right = node.Parent().RightChild()
			}

			if right.LeftChild().IsBlack() && right.RightChild().IsBlack() {
				right.SetColor(RED)
				node = node.Parent()
			} else {
				if right.RightChild().IsBlack() {
					right.LeftChild().SetColor(BLACK)
					right.SetColor(RED)
					tree.rightRotate(right)
					right = node.Parent().RightChild()
				}
				right.CopyColorFrom(node.Parent())
				node.Parent().SetColor(BLACK)
				right.RightChild().SetColor(BLACK)
				tree.leftRotate(node.Parent())
				node = tree.Root
			}
		} else {
			left := node.Parent().LeftChild()

			if left.IsRed() {
				left.SetColor(BLACK)
				node.Parent().SetColor(RED)
				tree.rightRotate(node.Parent())
				left = node.Parent().LeftChild()
			}
			if left.LeftChild().IsBlack() && left.RightChild().IsBlack() {
				left.SetColor(RED)
				node = node.Parent()
			} else {
				if left.LeftChild().IsBlack() {
					left.RightChild().SetColor(BLACK)
					left.SetColor(RED)
					tree.leftRotate(left)
					left = node.Parent().LeftChild()
				}
				left.CopyColorFrom(node.Parent())
				node.Parent().SetColor(BLACK)
				left.LeftChild().SetColor(BLACK)
				tree.rightRotate(node.Parent())
				node = tree.Root
			}
		}
	}

	node.SetColor(BLACK)
}

func (tree *Tree[K, V]) rebalance(node *Node[K, V]) {
	for node != tree.Root && node.Parent().IsRed() {
		if node.Parent() == node.Grandparent().LeftChild() {
			right := node.Grandparent().RightChild()

			if right.IsRed() {
				right.SetColor(BLACK)
				node.Parent().SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				node = node.Grandparent()
			} else {
				if node == node.Parent().RightChild() {
					node = node.Parent()
					tree.leftRotate(node)
				}
				node.Parent().SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				tree.rightRotate(node.Grandparent())
			}
		} else {
			left := node.Grandparent().LeftChild()

			if left.IsRed() {
				left.SetColor(BLACK)
				node.Parent().SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				node = node.Grandparent()
			} else {
				if node == node.Parent().LeftChild() {
					node = node.Parent()
					tree.rightRotate(node)
				}
				node.Parent().SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				tree.leftRotate(node.Grandparent())
			}
		}
	}
	tree.Root.SetColor(BLACK)
}

// Insert do an insertion to the tree.
func (tree *Tree[K, V]) Insert(key K, value V) (inserted *Node[K, V]) {
	if tree.Root.IsNil() {
		tree.Root = NewNode(tree.sentinel, key, value)
		inserted = tree.Root
		tree.Root.SetParent(nil)
	} else {
		// in redblack tree, the insert process can be different for same key.
		inserted = NewNode(tree.sentinel, key, value)
		indirectInsert, parent := tree.InsertLookupFn(tree.Root, key)
		*indirectInsert = inserted
		inserted.SetParent(parent)
	}
	tree.rebalance(inserted)
	tree.incr()
	return
}

// Get specific key from the tree
func (tree *Tree[K, V]) Get(key K) *Node[K, V] {
	return tree.LookupFn(tree.Root, key)
}

// Remove specific key from the tree
func (tree *Tree[K, V]) Remove(key K) bool {
	node := tree.LookupFn(tree.Root, key)
	if node == nil {
		return false
	}
	node.RemoveFrom(tree)
	return true
}

func output[K cmp.Ordered, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if !node.Right.IsNil() {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if !node.Left.IsNil() {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

// String returns a string representation of container
func (tree *Tree[K, V]) String() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

// Empty returns true if tree does not contain any nodes
func (tree *Tree[K, V]) Empty() bool {
	return tree.size == 0
}

// Min returns the minimum(leftmost) node of the tree.
//
// If the result is "nil" node, it will return Go's nil pointer instead.
func (tree *Tree[K, V]) Min() *Node[K, V] {
	return tree.Root.Min()
}

// Max returns the maximum(rightmost) node of the tree.
//
// If the result is "nil" node, it will return Go's nil pointer instead.
func (tree *Tree[K, V]) Max() *Node[K, V] {
	return tree.Root.Max()
}
