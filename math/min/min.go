package min

import "golang.org/x/exp/constraints"

// Int is a function that returns the minimum of all integers given as arguments.
func Int[T constraints.Integer](values ...T) T {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
