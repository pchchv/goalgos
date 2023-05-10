// Finding Mode Value In an Array.

package math

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// ErrEmptySlice is the error returned by functions in math package when
// an empty slice is provided to it as argument when the function expects
// a non-empty slice.
var ErrEmptySlice = errors.New("empty slice provided")

func Mode[T constraints.Integer](numbers []T) (T, error) {
	var mode T
	count := 0
	countMap := make(map[T]int)

	if len(numbers) == 0 {
		return 0, ErrEmptySlice
	}

	for _, number := range numbers {
		countMap[number]++
	}

	for k, v := range countMap {
		if v > count {
			count = v
			mode = k
		}
	}

	return mode, nil

}
