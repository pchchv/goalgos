// A simple implementation of Modular Inverse
// Reference: https://en.wikipedia.org/wiki/Modular_multiplicative_inverse

package modular

import (
	"errors"

	"github.com/pchchv/goalgos/math/gcd"
)

var ErrorInverse = errors.New("no Modular Inverse exists")

// Inverse Modular function
func Inverse(a, m int64) (int64, error) {
	gcd, x, _ := gcd.Extended(a, m)
	if gcd != 1 || m == 0 {
		return 0, ErrorInverse
	}
	// necessary because Go uses an architecture-specific
	// instruction for the % operator
	return ((m + (x % m)) % m), nil
}
