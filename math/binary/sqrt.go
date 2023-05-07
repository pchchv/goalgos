// Square root calculation.
// Calculating the square root using binary operations and
// a magic number 0x5f3759df.
// Reference: https://en.wikipedia.org/wiki/Fast_inverse_square_root

package binary

import (
	"math"
)

const threeHalves = 1.5

func Sqrt(n float32) float32 {
	var half, y float32
	half = n * 0.5
	z := math.Float32bits(n)
	z = 0x5f3759df - (z >> 1) // floating point bit level hacking
	y = math.Float32frombits(z)
	y = y * (threeHalves - (half * y * y)) // Newton's approximation
	return 1 / y
}
