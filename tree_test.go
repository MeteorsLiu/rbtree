package rbtree

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := NewTree[int, int]()

	node := []*Node[int, int]{}
	for i := 0; i < 10; i++ {
		node = append(node, tree.Insert(1, i))
	}
	// │               ┌── 1: 9: RED
	// │           ┌── 1: 8: BLACK
	// │       ┌── 1: 7: RED
	// │       │   └── 1: 6: BLACK
	// │   ┌── 1: 5: BLACK
	// │   │   └── 1: 4: BLACK
	// └── 1: 3: BLACK
	// 	   │   ┌── 1: 2: BLACK
	// 	   └── 1: 1: BLACK
	// 		   └── 1: 0: BLACK
	t.Log(tree.String())

	// left child is nil
	node[8].RemoveFrom(tree)

	// no child
	// node[3].RemoveFrom(tree)
	node[9].RemoveFrom(tree)

	// right child is nil
	node[7].RemoveFrom(tree)

	// two children
	node[1].RemoveFrom(tree)

	// │       ┌── 1: 6: BLACK
	// │   ┌── 1: 5: RED
	// │   │   └── 1: 4: BLACK
	// └── 1: 3: BLACK
	// 	   └── 1: 2: BLACK
	// 		   └── 1: 0: RED
	t.Log(tree.String())

}

func TestRedBlackTreePut10000(t *testing.T) {
	tree := NewTree[int, string]()

	for i := 0; i < 100000; i++ {
		tree.Insert(i, strconv.Itoa(i))
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(tree)
		}
	}()

	for i := 0; i < 150000; i++ {
		node := tree.Min()
		if node == nil {
			return
		}
		if node.Value != strconv.Itoa(i) || node.Key != i {
			t.Errorf("unexpected: want: %v  got: %v %v", i, node.Key, node.Value)
		}
		node.RemoveFrom(tree)
	}

	if !tree.Empty() {
		t.Error("unexpected: tree is not empty")
	}
	for i := 0; i < 100000; i++ {
		tree.Insert(i, strconv.Itoa(i))
	}
	if tree.Empty() {
		t.Error("unexpected: tree is empty")
	}
	for i := 0; i < 150000; i++ {
		node := tree.Min()
		if node == nil {
			return
		}
		if node.Value != strconv.Itoa(i) || node.Key != i {
			t.Errorf("unexpected: want: %v  got: %v %v", i, node.Key, node.Value)
		}
		node.RemoveFrom(tree)
	}
	if !tree.Empty() {
		t.Error("unexpected: tree is not empty")
	}
}

func TestRedBlackTreeGet(t *testing.T) {

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
		log.Println(node)
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
