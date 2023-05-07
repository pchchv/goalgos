// Reflected binary code (RBC).
// The reflected binary code (RBC), also known just as reflected binary (RB) or
// Gray code after Frank Gray, is an ordering of the binary numeral system such that
// two successive values differ in only one bit (binary digit).
// Reference: https://en.wikipedia.org/wiki/Gray_code

package binary

// SequenceGrayCode The function generates an "Gray code" sequence of length n
func SequenceGrayCode(n uint) []uint {
	result := make([]uint, 0)
	var i uint
	for i = 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	return result
}
