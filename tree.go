package rbtree

import (
	"cmp"
	"fmt"
)

type Tree[K cmp.Ordered, V any] struct {
	Root     *Node[K, V]
	size     int
	LookupFn LookupFn[K, V]
}

type Options[K cmp.Ordered, V any] func(*Tree[K, V])

func NewTree[K cmp.Ordered, V any](opts ...Options[K, V]) *Tree[K, V] {
	t := &Tree[K, V]{
		LookupFn: defaultLookup[K, V],
	}

	for _, o := range opts {
		o(t)
	}

	return t
}

func (tree *Tree[K, V]) leftRotate(node *Node[K, V]) {
	right := node.RightChild()
	node.Right = right.LeftChild()

	if right.LeftChild() != nil {
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

	if left.RightChild() != nil {
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
					tree.rightRotate(left)
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

func (tree *Tree[K, V]) Insert(key K, value V) (inserted *Node[K, V]) {
	if tree.Root == nil {
		tree.Root = NewNode(key, value)
		inserted = tree.Root
	} else {
		// in redblack tree, the insert process can be different for same key.
		inserted = NewNode(key, value)
		indirectInsert, parent := tree.LookupFn(tree.Root, key)
		*indirectInsert = inserted

		inserted.SetParent(parent)
	}
	tree.rebalance(inserted)
	tree.size++
	return
}

func (node *Node[K, V]) String() string {
	return fmt.Sprintf("%v: %v: %s", node.Key, node.Value, node.Color())
}

func output[K cmp.Ordered, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if node.Right != nil {
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
	if node.Left != nil {
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

func (tree *Tree[K, V]) Min() *Node[K, V] {
	return tree.Root.Min()
}

func (tree *Tree[K, V]) Max() *Node[K, V] {
	return tree.Root.Max()
}

func (tree *Tree[K, V]) replaceNode(old *Node[K, V], new *Node[K, V]) {
	if old.Parent() == nil {
		tree.Root = new
	} else {
		if old == old.Parent().LeftChild() {
			old.Parent().SetLeftChild(new)
		} else {
			old.Parent().SetRightChild(new)
		}
	}
	if new != nil {
		new.SetParent(old.Parent())
	}
}
