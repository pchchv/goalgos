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
