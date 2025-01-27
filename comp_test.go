package rbtree

import "testing"

func TestDefaultLookupGetDup(t *testing.T) {
	tree := NewTree[int, int]()

	expect := []int{1, 2, 3, 4, 5}

	for i := 1; i < 6; i++ {
		tree.Insert(1, i)
	}

	for _, want := range expect {
		node := tree.Get(1)
		if node == nil {
			t.Errorf("unexpected nil: %d", want)
		}
		if node.Value != want {
			t.Errorf("unexpected result: want %d got %d", want, node.Value)
		}
		node.RemoveFrom(tree)
	}
}

func TestDefaultLookupRemoveDup(t *testing.T) {
	tree := NewTree[int, int]()

	expect := []int{1, 2, 3, 4, 5}

	for i := 1; i < 6; i++ {
		tree.Insert(1, i)
	}
	pos := 0
	for _, want := range expect {
		if !tree.Remove(1) {
			t.Errorf("unexpected remove: %d", want)
		}
		node := tree.Get(1)
		if node == nil {
			t.Log(want)
			break
		}
		pos++

		if node.Value != expect[pos] {
			t.Errorf("unexpected result: want %d got %d", want, node.Value)
		}
	}

}

func TestDefaultLookupInsertDup(t *testing.T) {
	tree := NewTree[int, int]()

	expect := []int{1, 1, 1, 2, 3, 4, 5}

	for i, j := range expect {
		tree.Insert(j, i)
	}

	for i, want := range expect {
		node := tree.Get(want)
		if node == nil {
			t.Errorf("unexpected nil: %d", want)
		}
		if node.Value != i {
			t.Errorf("unexpected result: want %d got %d", want, node.Value)
		}
		node.RemoveFrom(tree)
	}

}
