// Find the median from a set of values.

package math

import (
	"github.com/pchchv/goalgos/sort"
	"golang.org/x/exp/constraints"
)

func Median[T constraints.Float](values []T) float64 {
	sort.Bubble(values)

	l := len(values)

	switch {
	case l == 0:
		return 0

	case l%2 == 0:
		return float64((values[l/2-1] + values[l/2]) / 2)

	default:
		return float64(values[l/2])
	}
}
