// Implementation of binary search algorithm.
// Description of the algorithm:
// https://en.wikipedia.org/wiki/Binary_search_algorithm
package search

// Binary search for a target in a sorted array by repeatedly dividing the
// array in half and comparing the midpoint to the target.
// This function uses a recursive call to itself.
// If the target is found, the target index is returned.
// Otherwise, the function returns -1 and ErrNotFound.
func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}

	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

// BinaryIterative searches for targets in a sorted array by repeatedly dividing the
// array in half and comparing the midpoint to the target.
// Unlike Binary, this function uses an iterative method rather than a recursive method.
// If the target is found, its index is returned.
// Otherwise, the function returns -1 and ErrNotFound.
func BinaryIterative(array []int, target int) (int, error) {
	var mid int
	startIndex := 0
	endIndex := len(array) - 1

	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}

	return -1, ErrNotFound
}

// LowerBound returns the index of the first element in the range [0, len(array)-1]
// that is not less than (i.e. greater than or equal to) target.
// Returns -1 and ErrNotFound if such element is not found.
func LowerBound(array []int, target int) (int, error) {
	var mid int
	startIndex := 0
	endIndex := len(array) - 1

	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] < target {
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}

	// when target greater than every element in array,
	// startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}

	return startIndex, nil
}

// UpperBound returns the index of the first element in
// the range [lowIndex, len(array)-1], which is larger than target.
// Returns -1 and ErrNotFound if such element is not found.
func UpperBound(array []int, target int) (int, error) {
	var mid int
	startIndex := 0
	endIndex := len(array) - 1

	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else {
			startIndex = mid + 1
		}
	}

	// when target greater or equal than every element in array,
	// startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}

	return startIndex, nil
}
