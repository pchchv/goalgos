// Absolute value.
// In mathematics, the absolute value or modulus of a real number x,
// denoted |x|, is the non-negative value of x without regard to its sign.
// Reference: https://en.wikipedia.org/wiki/Average#Arithmetic_mean

package math

// Abs returns absolute value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
