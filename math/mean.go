package math

import "golang.org/x/exp/constraints"

func Mean[T constraints.Float](values []T) float64 {
	if len(values) == 0 {
		return 0
	}

	var summation float64 = 0

	for _, singleValue := range values {
		summation += float64(singleValue)
	}

	return summation / float64(len(values))
}
