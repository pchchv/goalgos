package tree_test

import (
	"fmt"
	"testing"

	bt "github.com/pchchv/goalgos/structure/tree"
	"golang.org/x/exp/constraints"
)

type TestTree[T constraints.Ordered] interface {
	Push(...T)
	Delete(T) bool
	Get(T) (bt.Node[T], bool)
	Empty() bool
	Has(T) bool
	Depth() int
	Max() (T, bool)
	Min() (T, bool)
	Predecessor(T) (T, bool)
	Successor(T) (T, bool)
	PreOrder() []T
	InOrder() []T
	PostOrder() []T
	LevelOrder() []T
	AccessNodesByLayer() [][]T
}

var (
	// BinarySearch, AVL and RB have completed the `TestTree` interface.
	_ TestTree[int] = (*bt.BinarySearch[int])(nil)
	_ TestTree[int] = (*bt.AVL[int])(nil)

	tree TestTree[int]
)

func TestBinarySearch(t *testing.T) {
	tree = bt.NewBinarySearch[int]()

	if tree.Empty() {
		t.Log("Binary Search Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
}

func TestAVL(t *testing.T) {
	tree = bt.NewAVL[int]()

	if tree.Empty() {
		t.Log("AVL Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
}
