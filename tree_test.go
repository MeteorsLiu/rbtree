package rbtree

import (
	"cmp"
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
	if tree.Leftmost() == nil || tree.Leftmost().Key != 0 {
		t.Error("leftmost", tree.Leftmost())
	}
	tree.Print()
	for i := 0; i < 10; i++ {
		n := tree.Search(i)
		if n == nil || n.Data != r[i] {
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

	// remove again
	tree.Delete(rn[rm])
	//tree.Print()

	n := tree.Search(rm)
	if n != nil {
		t.Error("error deleted again", rm, n)
	}

	tree.InsertNode(rn[rm])

	n = tree.Search(rm)
	if n != rn[rm] || n == nil {
		t.Error("error", rm, n)
	}

	fmt.Println("After insert deleted node: ")
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
func testCustomCmp(a, b int) int {
	if a > b {
		return GREATER
	} else if a < b {
		return LESS
	} else {
		return EQUAL
	}
}

func BenchmarkTreeCustomCMP(b *testing.B) {
	tree := NewRBTree[int, int](testCustomCmp)
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

func BenchmarkCmp(b *testing.B) {
	c, d := rand.Int(), rand.Int()
	b.Run("GoCompare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmp.Compare(c, d)
		}
	})

}
