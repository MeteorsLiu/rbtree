package rbtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type RBTree[K constraints.Ordered, T any] struct {
	// nilNode stands for nil elements in RBTree, which is black.
	nilNode *RBNode[K, T]
	root    *RBNode[K, T]
}

func NewRBTree[K constraints.Ordered, T any]() *RBTree[K, T] {
	var zero K
	nilNode := NewRBNode[K, T](nil, nil, nil, zero)
	nilNode.SetColor(BLACK)
	return &RBTree[K, T]{nilNode: nilNode, root: nilNode}
}

// Insert a key and data into the RBTree, if the key exists, return the node and wether it's inserted or not.
func (t *RBTree[K, T]) Insert(key K, data T) (*RBNode[K, T], bool) {
	node := t.root

	if node == t.nilNode {
		t.root = NewRBNode[K, T](t.nilNode, t.nilNode, t.nilNode, key, data)
		t.root.SetColor(BLACK)
		return t.root, true
	}

	for {
		if key < node.key {
			if node.left == t.nilNode {
				node.left = NewRBNode[K, T](node, t.nilNode, t.nilNode, key, data)
				node = node.left
				break
			}
			node = node.left
		} else if key > node.key {
			if node.right == t.nilNode {
				node.right = NewRBNode[K, T](node, t.nilNode, t.nilNode, key, data)
				node = node.right
				break
			}
			node = node.right
		} else {
			// hey, see what we found, a duplicate key
			// cannot insert so return here.
			return node, false
		}
	}
	oldNode := node

	for node != t.root && node.parent.IsRed() {
		if node.parent.IsLeftChild() {
			uncle := node.Grandparent().Right()

			if uncle.IsRed() {
				node.parent.SetColor(BLACK)
				uncle.SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				node = node.Grandparent()
			} else {
				if node.IsRightChild() {
					node = node.parent
					t.RotateLeft(node)
				}
				node.parent.SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				t.RotateRight(node.Grandparent())
			}
		} else {
			uncle := node.Grandparent().Left()

			if uncle.IsRed() {
				node.parent.SetColor(BLACK)
				uncle.SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				node = node.Grandparent()
			} else {
				if node.IsLeftChild() {
					node = node.parent
					t.RotateRight(node)
				}
				node.parent.SetColor(BLACK)
				node.Grandparent().SetColor(RED)
				t.RotateLeft(node.Grandparent())
			}
		}
	}

	t.root.SetColor(BLACK)
	return oldNode, true
}

func (t *RBTree[K, T]) Search(key K) *RBNode[K, T] {
	node := t.root

	for node != t.nilNode && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	if node == t.nilNode {
		return nil
	}

	return node
}

// 删除太难了，抄袭了Nginx的代码实现
func (t *RBTree[K, T]) Delete(n *RBNode[K, T]) {
	if n == t.nilNode || n == nil {
		return
	}
	var ptr, subst *RBNode[K, T]

	if n.left == t.nilNode {
		subst = n
		ptr = n.right
	} else if n.right == t.nilNode {
		subst = n
		ptr = n.left
	} else {
		subst = t.Min(n)
		ptr = n.right
	}

	if subst == t.root {
		t.root = ptr
		t.root.SetColor(BLACK)
		return
	}

	substIsRed := subst.IsRed()

	if subst.IsLeftChild() {
		subst.parent.left = ptr
	} else {
		subst.parent.right = ptr
	}

	if subst == n {
		ptr.parent = subst.parent
	} else {
		if subst.parent == n {
			ptr.parent = subst
		} else {
			ptr.parent = subst.parent
		}
		subst.left = n.left
		subst.right = n.right
		subst.parent = n.parent
		subst.SetColor(n.Color())

		if n == t.root {
			t.root = subst
		} else {
			if n.IsLeftChild() {
				n.parent.left = subst
			} else {
				n.parent.right = subst
			}
		}

		if subst.left != t.nilNode {
			subst.left.parent = subst
		}

		if subst.right != t.nilNode {
			subst.right.parent = subst
		}
	}

	if substIsRed {
		return
	}

	for ptr != t.root && ptr.IsBlack() {
		if ptr.IsLeftChild() {
			sibling := ptr.parent.right

			if sibling.IsRed() {
				sibling.SetColor(BLACK)
				ptr.parent.SetColor(RED)
				t.RotateLeft(ptr.parent)
				sibling = ptr.parent.right
			}

			if sibling.left.IsBlack() && sibling.right.IsBlack() {
				sibling.SetColor(RED)
				ptr = ptr.parent
			} else {
				if sibling.right.IsBlack() {
					sibling.left.SetColor(BLACK)
					sibling.SetColor(RED)
					t.RotateRight(sibling)
					sibling = ptr.parent.right
				}

				sibling.SetColor(ptr.parent.Color())
				ptr.parent.SetColor(BLACK)
				sibling.right.SetColor(BLACK)
				t.RotateLeft(ptr.parent)
				ptr = t.root
			}
		} else {
			sibling := n.parent.left
			if sibling.IsRed() {
				sibling.SetColor(BLACK)
				ptr.parent.SetColor(RED)
				t.RotateRight(ptr.parent)
				sibling = ptr.parent.left
			}

			if sibling.left.IsBlack() && sibling.right.IsBlack() {
				sibling.SetColor(RED)
				ptr = ptr.parent
			} else {
				if sibling.left.IsBlack() {
					sibling.right.SetColor(BLACK)
					sibling.SetColor(RED)
					t.RotateLeft(sibling)
					sibling = ptr.parent.left
				}

				sibling.SetColor(ptr.parent.Color())
				ptr.parent.SetColor(BLACK)
				sibling.left.SetColor(BLACK)
				t.RotateRight(ptr.parent)
				ptr = t.root
			}
		}
	}

	ptr.SetColor(BLACK)
}

func (t *RBTree[K, T]) Leftmost() *RBNode[K, T] {
	if left := t.root.left; t.root != t.nilNode && left != t.nilNode {
		return left
	}
	return nil
}

func (t *RBTree[K, T]) walk(n *RBNode[K, T], f func(int, bool, *RBNode[K, T]), depth int) {
	if n == nil {
		return
	}
	f(depth, n == t.nilNode, n)
	t.walk(n.left, f, depth+1)
	t.walk(n.right, f, depth+1)
}

func (t *RBTree[K, T]) Walk(f func(int, bool, *RBNode[K, T])) {
	t.walk(t.root, f, 0)
}

func (t *RBTree[K, T]) print(n *RBNode[K, T], space int) {
	if n == nil {
		return
	}
	space += 10
	t.print(n.right, space)

	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}

	if n == t.nilNode {
		fmt.Printf("%s: Nil\n", n.Color())
	} else {
		fmt.Printf("%s: %v => %v\n", n.Color(), n.Key(), n.Data())
	}
	t.print(n.left, space)
}

func (t *RBTree[K, T]) Print() {
	t.print(t.root, 0)
}
