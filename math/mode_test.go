package math_test

import (
	"errors"
	"testing"

	"github.com/pchchv/goalgos/math"
	"golang.org/x/exp/constraints"
)

type testCase[T constraints.Integer] struct {
	name    string
	numbers []T
	mode    T
	err     error
}

func testModeFramework[T constraints.Integer](t *testing.T, testCases []testCase[T]) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnMode, err := math.Mode(test.numbers)
			if returnMode != test.mode {
				t.Errorf("\n Failed test %s,\n Numbers: %v,\n Correct Mode: %v,\n Returned Mode: %v\n",
					test.name, test.numbers, test.mode, returnMode)
			}
			if !errors.Is(err, test.err) {
				t.Errorf("\n Failed test %s,\n Numbers: %v,\n Correct Error: %v,\n Returned Error: %v\n",
					test.name, test.numbers, test.err, err)
			}
		})
	}
}
func TestMode(t *testing.T) {
	// test cases for integer values
	intTestCases := []testCase[int]{
		{
			name:    "An array of positive whole numbers",
			numbers: []int{10, 52, 10, 92, 10, 75, 60, 10, 44, 29},
			mode:    10,
			err:     nil,
		},
		{
			name:    "An array of negative whole numbers",
			numbers: []int{-19, -12, -74, -19, -22, -56, -19, -19, -68, -93},
			mode:    -19,
			err:     nil,
		},
		{
			name:    "An array of positive & negative whole numbers",
			numbers: []int{18, -28, 33, -28, 2, 39, 48, -49, -87, 78, -28},
			mode:    -28,
			err:     nil,
		},
		{
			name:    "If array has no value",
			numbers: []int{},
			mode:    0,
			err:     math.ErrEmptySlice,
		},
	}
	testModeFramework(t, intTestCases)
}
