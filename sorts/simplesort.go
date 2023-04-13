// Implementation of a simple sorting algorithm
// Reference: https://arxiv.org/abs/2110.01111v1
// see sort_test.go for a test implementation, test function TestSimple and TestImprovedSimple

package sorts

import "golang.org/x/exp/constraints"

func Simple[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// ImprovedSimple is an improvement to SimpleSort by skipping the unnecessary first-to-last comparison.
// This improved version is more like an implementation of insertion sorting.
func ImprovedSimple[T constraints.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
