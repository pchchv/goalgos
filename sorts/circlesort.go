// Implementation of circle sort algorithm.
// Reference: https://www.geeksforgeeks.org/circle-sort/
package sorts

import (
	"golang.org/x/exp/constraints"
)

func circle[T constraints.Ordered](a []T, lo, hi, swaps int) int {
	if hi < 2 {
		return swaps
	}

	if lo == hi {
		return swaps
	}

	high, low := hi, lo
	mid := (hi - lo) / 2

	for lo < hi {
		if a[lo] > a[hi] {
			a[lo], a[hi] = a[hi], a[lo]
			swaps++
		}
		lo++
		hi--
	}

	if lo == hi {
		if a[lo] > a[hi+1] {
			a[lo], a[hi+1] = a[hi+1], a[lo]
			swaps++
		}
	}

	swaps = circle(a, low, low+mid, swaps)
	swaps = circle(a, low+mid+1, high, swaps)

	return swaps
}

func Circle[T constraints.Ordered](items []T) []T {
	for circle(items, 0, len(items)-1, 0) != 0 {
	}

	return items
}
