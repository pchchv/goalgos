// Package diffiehellman implements Diffie-Hellman Key Exchange Algorithm
// Reference: https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
package diffiehellman

const (
	generator         = 3
	primeNumber int64 = 6700417 // prime number discovered by Leonhard Euler
)

// GenerateShareKey generates a key using the client's private key,
// generator and primeNumber, this key can be made public
// shareKey = (g^key)%primeNumber.
func GenerateShareKey(prvKey int64) int64 {
	return modularExponentiation(generator, prvKey, primeNumber)
}

// GenerateMutualKey generates a mutual key that can only be used by Alice and Bob
// mutualKey = (shareKey^prvKey)%primeNumber
func GenerateMutualKey(prvKey, shareKey int64) int64 {
	return modularExponentiation(shareKey, prvKey, primeNumber)
}

// modularExponentiation runs in O(log(n)), where n = e,
// uses exponentiation over the square to speed up the process
// r = (b^e)%mod
func modularExponentiation(b, e, mod int64) int64 {
	if mod == 1 {
		return 0
	}

	var r int64 = 1
	b = b % mod

	for e > 0 {
		if e&1 == 1 {
			r = (r * b) % mod
		}
		e = e >> 1
		b = (b * b) % mod
	}

	return r
}
