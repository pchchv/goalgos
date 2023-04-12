package search

type searchTest struct {
	data          []int
	key           int
	expected      int
	expectedError error
	name          string
}

// Note that they are immutable, so they are common to all search tests.
var searchTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},

	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, ErrNotFound, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},

	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}

var lowerBoundTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -25, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 2, 2, 2, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, 0, nil, "Sanity"},

	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},

	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}

var upperBoundTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -25, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 5, nil, "Sanity"},
	{[]int{1, 2, 2, 2, 2, 6, 7, 8, 9, 10}, 2, 5, nil, "Sanity"},

	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, -1, ErrNotFound, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, -1, ErrNotFound, "Sanity"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},

	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}
