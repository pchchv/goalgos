// Cosine calculation.
// Reference: https://en.wikipedia.org/wiki/Sine_and_cosine

package math

import "math"

// Cos returns the cosine of the radian argument x.
func Cos(x float64) float64 {
	tp := 1.0 / (2.0 * math.Pi)
	x *= tp
	x -= 0.25 + math.Floor(x+0.25)
	x *= 16.0 * (math.Abs(x) - 0.5)
	x += 0.225 * x * (math.Abs(x) - 1.0) // extra precision
	return x
}
