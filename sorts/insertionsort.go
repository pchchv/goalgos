// Implementation of insertion sort algorithm.
// Reference: https://en.wikipedia.org/wiki/Insertion_sort
package sorts

import "golang.org/x/exp/constraints"

func Insertion[T constraints.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex

		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}

	return arr
}
