// Lerp or Linear interpolation
// Reference: https://en.wikipedia.org/wiki/Linear_interpolation
package math

// This function will return new value in 't' percentage  between 'v0' and 'v1'
func Lerp(v0, v1, t float64) float64 {
	return (1-t)*v0 + t*v1
}
