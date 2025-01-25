package rbtree

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := NewTree[int, int]()

	for i := 0; i < 10; i++ {
		tree.Insert(1, i)

	}

	t.Log(tree.String())

	tree.Min().RemoveFrom(tree)

	t.Log(tree.String())

}

func TestRedBlackTreePut10000(t *testing.T) {
	tree := NewTree[int, string]()

	for i := 0; i < 100000; i++ {
		tree.Insert(i, strconv.Itoa(i))
	}

	for i := 0; i < 100000; i++ {
		node := tree.Min()
		if node == nil {
			return
		}
		if node.Value != strconv.Itoa(i) || node.Key != i {
			t.Errorf("unexpected: want: %v  got: %v %v", i, node.Key, node.Value)
		}
		node.RemoveFrom(tree)
	}
}

func TestRedBlackTreePut(t *testing.T) {
	tree := NewTree[int, string]()
	tree.Insert(5, "e")
	tree.Insert(6, "f")
	tree.Insert(7, "g")
	tree.Insert(3, "c")
	tree.Insert(4, "d")
	tree.Insert(1, "x")
	tree.Insert(2, "b")
	tree.Insert(1, "a") //overwrite

	tests1 := [][]interface{}{
		{1, "x", true},
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		node := tree.Min()
		if node == nil {
			return
		}
		if node.Value != test[1] || node.Key != test[0] {
			t.Errorf("unexpected: want: %v %v got: %v %v", test[0], test[1], node.Key, node.Value)
		}
		node.RemoveFrom(tree)
	}
}

func benchmarkPutRemove(b *testing.B, tree *Tree[int, struct{}]) {
	for n := 0; n < b.N; n++ {
		tree.Insert(rand.Int(), struct{}{})
	}

	for n := 0; n < b.N; n++ {
		tree.Min().RemoveFrom(tree)
	}
}

func BenchmarkRedBlackTreePut(b *testing.B) {
	tree := NewTree[int, struct{}]()

	benchmarkPutRemove(b, tree)
}
