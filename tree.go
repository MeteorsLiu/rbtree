package rbtree

import (
	"cmp"
	"fmt"
)

type RBTree[K cmp.Ordered, T any] struct {
	// nilNode stands for nil elements in RBTree, which is black.
	nilNode *RBNode[K, T]
	root    *RBNode[K, T]
	cmp     Compare[K]
}

func NewRBTree[K cmp.Ordered, T any](customCompare ...Compare[K]) *RBTree[K, T] {
	var zero K
	nilNode := NewRBNode[K, T](nil, nil, nil, zero)
	nilNode.setColor(BLACK)

	if len(customCompare) > 0 {
		return &RBTree[K, T]{nilNode: nilNode, root: nilNode, cmp: customCompare[0]}
	}
	return &RBTree[K, T]{nilNode: nilNode, root: nilNode}
}

func (t *RBTree[K, T]) insertBalance(node *RBNode[K, T]) {
	for node != t.root && node.parent.IsRed() {
		if node.parent.IsLeftChild() {
			uncle := node.Grandparent().Right()

			if uncle.IsRed() {
				node.parent.setColor(BLACK)
				uncle.setColor(BLACK)
				node.Grandparent().setColor(RED)
				node = node.Grandparent()
			} else {
				if node.IsRightChild() {
					node = node.parent
					t.RotateLeft(node)
				}
				node.parent.setColor(BLACK)
				node.Grandparent().setColor(RED)
				t.RotateRight(node.Grandparent())
			}
		} else {
			uncle := node.Grandparent().Left()

			if uncle.IsRed() {
				node.parent.setColor(BLACK)
				uncle.setColor(BLACK)
				node.Grandparent().setColor(RED)
				node = node.Grandparent()
			} else {
				if node.IsLeftChild() {
					node = node.parent
					t.RotateRight(node)
				}
				node.parent.setColor(BLACK)
				node.Grandparent().setColor(RED)
				t.RotateLeft(node.Grandparent())
			}
		}
	}
	t.root.setColor(BLACK)
}

// Insert a key and data into the RBTree, if the key exists, return the node and wether it's inserted or not.
func (t *RBTree[K, T]) Insert(key K, data T) (*RBNode[K, T], bool) {
	node := t.root

	if node == t.nilNode {
		t.root = NewRBNode[K, T](t.nilNode, t.nilNode, t.nilNode, key, data)
		t.root.setColor(BLACK)
		return t.root, true
	}
	var r int
	for {
		if t.cmp != nil {
			r = t.cmp(key, node.Key)
		} else {
			r = cmp.Compare(key, node.Key)
		}
		if r == LESS {
			if node.left == t.nilNode {
				node.left = NewRBNode[K, T](node, t.nilNode, t.nilNode, key, data)
				node = node.left
				break
			}
			node = node.left
		} else if r == GREATER {
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
	t.insertBalance(node)
	return oldNode, true
}

func (t *RBTree[K, T]) InsertNode(n *RBNode[K, T]) (*RBNode[K, T], bool) {
	node := t.root

	if node == t.nilNode {
		n.reset(t.nilNode, t.nilNode, t.nilNode, BLACK)
		t.root = n
		return t.root, true
	}

	var r int
	for {
		if t.cmp != nil {
			r = t.cmp(n.Key, node.Key)
		} else {
			r = cmp.Compare(n.Key, node.Key)
		}
		if r == LESS {
			if node.left == t.nilNode {
				n.reset(node, t.nilNode, t.nilNode, RED)
				node.left = n
				node = node.left
				break
			}
			node = node.left
		} else if r == GREATER {
			if node.right == t.nilNode {
				n.reset(node, t.nilNode, t.nilNode, RED)
				node.right = n
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
	t.insertBalance(node)
	return n, true
}

func (t *RBTree[K, T]) Search(key K) *RBNode[K, T] {
	node := t.root
	var r int
	for node != t.nilNode {
		if t.cmp != nil {
			r = t.cmp(key, node.Key)
		} else {
			r = cmp.Compare(key, node.Key)
		}
		if r == LESS {
			node = node.left
		} else if r == GREATER {
			node = node.right
		} else {
			break
		}
	}
	if node == t.nilNode {
		return nil
	}

	return node
}

func (t *RBTree[K, T]) Delete(n *RBNode[K, T]) bool {
	if n == t.nilNode || n == nil {
		return false
	}
	var r int
	node := t.root
	for node != t.nilNode {
		if t.cmp != nil {
			r = t.cmp(n.Key, node.Key)
		} else {
			r = cmp.Compare(n.Key, node.Key)
		}
		if r == LESS {
			node = node.left
		} else if r == GREATER {
			node = node.right
		} else {
			break
		}
	}

	if node == t.nilNode || node != n {
		return false
	}

	t.DeleteUnsafe(node)
	return true
}

// Unsafe delete, it will not check wether the n is in the tree
// if n is not in the tree, it will panic.
// this function is for performance when you can ensure n is in the tree,
// otherwise, use Delete() instead.
func (t *RBTree[K, T]) DeleteUnsafe(n *RBNode[K, T]) {
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
		subst = t.Min(n.right)
		ptr = subst.right
	}

	if subst == t.root {
		t.root = ptr
		t.root.setColor(BLACK)
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
		subst.setColor(n.Color())

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
				sibling.setColor(BLACK)
				ptr.parent.setColor(RED)
				t.RotateLeft(ptr.parent)
				sibling = ptr.parent.right
			}

			if sibling.left.IsBlack() && sibling.right.IsBlack() {
				sibling.setColor(RED)
				ptr = ptr.parent
			} else {
				if sibling.right.IsBlack() {
					sibling.left.setColor(BLACK)
					sibling.setColor(RED)
					t.RotateRight(sibling)
					sibling = ptr.parent.right
				}

				sibling.setColor(ptr.parent.Color())
				ptr.parent.setColor(BLACK)
				sibling.right.setColor(BLACK)
				t.RotateLeft(ptr.parent)
				ptr = t.root
			}
		} else {
			sibling := n.parent.left
			if sibling.IsRed() {
				sibling.setColor(BLACK)
				ptr.parent.setColor(RED)
				t.RotateRight(ptr.parent)
				sibling = ptr.parent.left
			}

			if sibling.left.IsBlack() && sibling.right.IsBlack() {
				sibling.setColor(RED)
				ptr = ptr.parent
			} else {
				if sibling.left.IsBlack() {
					sibling.right.setColor(BLACK)
					sibling.setColor(RED)
					t.RotateLeft(sibling)
					sibling = ptr.parent.left
				}

				sibling.setColor(ptr.parent.Color())
				ptr.parent.setColor(BLACK)
				sibling.left.setColor(BLACK)
				t.RotateRight(ptr.parent)
				ptr = t.root
			}
		}
	}

	ptr.setColor(BLACK)
}

func (t *RBTree[K, T]) Leftmost() *RBNode[K, T] {
	return t.Min(t.root)
}

func (t *RBTree[K, T]) Rightmost() *RBNode[K, T] {
	return t.Max(t.root)
}

func (t *RBTree[K, T]) Min(n *RBNode[K, T]) *RBNode[K, T] {
	for n.left != t.nilNode {
		n = n.left
	}
	if n == t.nilNode {
		return nil
	}
	return n
}

func (t *RBTree[K, T]) Max(n *RBNode[K, T]) *RBNode[K, T] {
	for n.right != t.nilNode {
		n = n.right
	}
	if n == t.nilNode {
		return nil
	}
	return n
}

func (t *RBTree[K, T]) Next(n *RBNode[K, T]) *RBNode[K, T] {
	if n.right != t.nilNode {
		return t.Min(n.right)
	}

	for n != t.root && n.IsRightChild() {
		n = n.parent
	}

	if n == t.nilNode {
		return nil
	}
	return n
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
		fmt.Printf("%s: %v => %v\n", n.Color(), n.Key, n.Data)
	}
	t.print(n.left, space)
}

func (t *RBTree[K, T]) Print() {
	t.print(t.root, 0)
}
