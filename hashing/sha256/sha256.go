// The sha256 cryptographic hash function as defined in the RFC6234 standard.
// Reference: https://en.wikipedia.org/wiki/SHA-2
// Reference: https://datatracker.ietf.org/doc/html/rfc6234
package sha256

import (
	"encoding/binary" // Used for interacting with uint at the byte level
)

// pad returns an extended version of the input message,
// e.g., the length of the extended message is a multiple of 512 bits.
// The padding methodology is as follows:
// A "1" bit is added at the end of the input message,
// followed by m "0" bits, so that the message length is 64 bits less than the length multiple of 512 bits.
// The remaining 64 bits are filled with the initial message length, represented as a 64-bit unsigned integer.
// For more information, see:
// https://datatracker.ietf.org/doc/html/rfc6234#section-4.1.
func pad(message []byte) []byte {
	L := make([]byte, 8)
	binary.BigEndian.PutUint64(L, uint64(len(message)*8))
	message = append(message, 0x80) // "1" bit followed by 7 "0" bits

	for (len(message)+8)%64 != 0 {
		message = append(message, 0x00) // 8 "0" bits
	}

	message = append(message, L...)
	return message
}
