// Checks if the given number is armstrong number or not.
// An Armstrong number is a n-digit number that is equal to the sum of each of its digits taken to the nth power.
// Reference: https://mathlair.allfunandgames.ca/armstrong.php

package armstrong

import (
	"math"
	"strconv"
)

func IsArmstrong(number int) bool {
	var rightMost int
	var sum int = 0
	var tempNum int = number

	// to get the number of digits in the number
	length := float64(len(strconv.Itoa(number)))

	// get the right most digit and break the loop once all digits are iterated
	for tempNum > 0 {
		rightMost = tempNum % 10
		sum += int(math.Pow(float64(rightMost), length))

		// update the input digit minus the processed rightMost
		tempNum /= 10
	}

	return number == sum
}
