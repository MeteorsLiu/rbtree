package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRbTree(t *testing.T) {
	tree := NewRBTree[int, int]()
	bst := NewBSTree[int, int]()
	r := []int{}
	rn := []*RBNode[int, int]{}
	for i := 0; i < 10; i++ {
		r = append(r, rand.Int())
		tree.Insert(i, r[i])
		bst.Insert(i, r[i])
	}
	if tree.Leftmost() == nil || tree.Leftmost().Key() != 0 {
		t.Error("leftmost", tree.Leftmost())
	}
	tree.Print()
	for i := 0; i < 10; i++ {
		n := tree.Search(i)
		if n == nil || n.Data() != r[i] {
			t.Error("error", i, n)
		}
		rn = append(rn, n)
		//t.Log(i, n.Color(), n.Parent(), n.IsLeftChild(), n.Key())
	}

	rm := rand.Intn(len(rn) - 1)
	t.Log("removed", rm)
	tree.Delete(rn[rm])

	for i := 0; i < 10; i++ {
		n := tree.Search(i)
		if i == rm && n != nil {
			t.Error("error", i, n)
		}
		//t.Log(i, n.Color(), n.Parent(), n.IsLeftChild(), n.Key())
	}
	fmt.Println("After deleted: ")
	tree.Print()
}

func BenchmarkTree(b *testing.B) {
	tree := NewRBTree[int, int]()
	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree.Insert(i, i)
		}
	})

	b.Run("Search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = tree.Search(i)
		}
	})
}

func BenchmarkMap(b *testing.B) {
	m := map[int]int{}
	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[i] = i
		}
	})

	b.Run("Search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[i]
		}
	})
}
