// Gives max of two integers
// details: implementation of finding the maximum of
// two numbers using only binary operations without using conditions.

package max

// Bitwise computes using bitwise operator the maximum of all the integer input and returns it
func Bitwise(a int, b int, base int) int {
	z := a - b
	i := (z >> base) & 1
	return a - (i * z)
}
