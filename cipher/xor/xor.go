// Package xor is an encryption algorithm that operates the exclusive disjunction(XOR)
// Reference: https://en.wikipedia.org/wiki/XOR_cipher
package xor

// Encrypt encrypts using Xor encryption after converting each character to a byte.
// The returned value may be unreadable because there is no guarantee that it is in
// the ASCII range When using another type such as string,
// []int or some other types, add statements to convert the type to []byte.
func Encrypt(key byte, plaintext []byte) (cipherText []byte) {
	for _, ch := range plaintext {
		cipherText = append(cipherText, key^ch)
	}
	return
}

// Decrypt decrypts with Xor encryption
func Decrypt(key byte, cipherText []byte) (plainText []byte) {
	for _, ch := range cipherText {
		plainText = append(plainText, key^ch)
	}
	return
}
