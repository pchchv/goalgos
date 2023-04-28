package tree_test

import (
	"testing"

	bt "github.com/pchchv/goalgos/structure/tree"
)

// Benchmark the comparisons between BST, AVL and RB Tree
const testNum = 10_000

func BenchmarkAVLTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}

}

func BenchmarkAVLTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkAVLTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkBSTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkBSTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkBSTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}
