// Counts the number of set bits in a number.
// For unsigned integer number N, return the number of bits set to 1
// Reference: https://en.wikipedia.org/wiki/Bit_numbering

package binary

// BitCounter - The function returns the number of set bits for an unsigned integer number
func BitCounter(n uint) int {
	counter := 0
	for n != 0 {
		if n&1 == 1 {
			counter++
		}
		n >>= 1
	}
	return counter
}
