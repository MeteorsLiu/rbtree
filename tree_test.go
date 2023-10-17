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
