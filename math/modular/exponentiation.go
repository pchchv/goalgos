// Implementation of Modular Exponentiation Algorithm
// A simple implementation of Modular Exponentiation - [Modular Exponenetation wiki](https://en.wikipedia.org/wiki/Modular_exponentiation)
package modular

import (
	"errors"
	"math"
)

var (
	// ErrorIntOverflow For asserting that the values do not overflow in Int64
	ErrorIntOverflow = errors.New("integer overflow")
	// ErrorNegativeExponent for asserting that the exponent we receive is positive
	ErrorNegativeExponent = errors.New("negative Exponent provided")
)

// Exponentiation returns base^exponent % mod
func Exponentiation(base, exponent, mod int64) (int64, error) {
	if mod == 1 {
		return 0, nil
	}

	if exponent < 0 {
		return -1, ErrorNegativeExponent
	}
	_, err := Multiply64BitInt(mod-1, mod-1)

	if err != nil {
		return -1, err
	}

	var result int64 = 1

	base = base % mod

	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		exponent = exponent >> 1
		base = (base * base) % mod
	}
	return result, nil
}

// Multiply64BitInt checking if the integer multiplication overflows
func Multiply64BitInt(left, right int64) (int64, error) {
	if math.Abs(float64(left)) > float64(math.MaxInt64)/math.Abs(float64(right)) {
		return 0, ErrorIntOverflow
	}

	return left * right, nil
}
