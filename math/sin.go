// Sine calculation.
// Referencd: https://en.wikipedia.org/wiki/Sine_and_cosine
package math

import "math"

// Sin returns the sine of the radian argument x.
func Sin(x float64) float64 {
	return Cos((math.Pi / 2) - x)
}
