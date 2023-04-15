package gcd

// Iterative is a faster iterative version of
// GcdRecursive that does not take up too much stack.
func Iterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
