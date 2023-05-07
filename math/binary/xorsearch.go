// Find a missing number in a sequence.
// Given an array A containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
// Reference: https://en.wikipedia.org/wiki/Exclusive_or

package binary

// XorSearchMissingNumber This function finds a missing number in a sequence
func XorSearchMissingNumber(a []int) int {
	n := len(a)
	result := len(a)
	for i := 0; i < n; i++ {
		result ^= i ^ a[i]
	}
	return result
}
