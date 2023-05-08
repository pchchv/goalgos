// Provides the functions
// - IsPerfectNumber which checks if the input is a perfect number,
// - SumOfProperDivisors which returns the sum of proper divisors of the input.
// A number is called perfect, if it is a sum of its proper divisors.
// Reference: https://en.wikipedia.org/wiki/Perfect_number,

package math

// Returns the sum of proper divisors of inNumber.
func SumOfProperDivisors(inNumber uint) uint {
	var res = uint(0)
	if inNumber > 1 {
		res = uint(1)
	}
	for curDivisor := uint(2); curDivisor*curDivisor <= inNumber; curDivisor++ {
		if inNumber%curDivisor == 0 {
			res += curDivisor
			if curDivisor*curDivisor != inNumber {
				res += inNumber / curDivisor
			}
		}
	}
	return res
}

// Checks if inNumber is a perfect number
func IsPerfectNumber(inNumber uint) bool {
	return inNumber > 0 && SumOfProperDivisors(inNumber) == inNumber
}
