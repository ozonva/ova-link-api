package utils

import "testing"

var chunkTestCases = []struct {
	caseName   string
	inputSlice []int
	inputSize  int
	expected   [][]int
}{
	{"empty slice", []int{}, 3, [][]int{}},
	{"one element, chunked equally", []int{1}, 1, [][]int{{1}}},
	{"several elements, size is more than length", []int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
	{"several elements, size = 0", []int{1, 2, 3}, 0, [][]int{}},
	{"several elements, size < 0", []int{1, 2, 3}, -1, [][]int{}},
	{"odd amount of elements, chunked equally", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	{"odd amount of elements, chunked unequally", []int{1, 2, 3, 4, 5, 6, 7}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7}}},
	{"even amount of elements, chunked equally", []int{1, 2, 3, 4, 5, 6}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}}},
	{"even amount of elements, chunked unequally", []int{1, 2, 3, 4, 5, 6}, 4, [][]int{{1, 2, 3, 4}, {5, 6}}},
}

func TestSliceChunk(t *testing.T) {
	for _, testCase := range chunkTestCases {
		result := SliceChunk(testCase.inputSlice, testCase.inputSize)

		if len(result) != len(testCase.expected) {
			t.Fatalf(`%v. Expected length: %v. Actual length: %v`, testCase.caseName, len(testCase.expected), len(result))
		}

		for i, chunk := range result {
			for j, element := range chunk {
				if element != testCase.expected[i][j] {
					t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
				}
			}
		}
	}
}

var filterByListTestCases = []struct {
	caseName   string
	inputSlice []int
	expected   []int
}{
	{"empty slice", []int{}, []int{}},
	{"one element, not in filter", []int{4}, []int{4}},
	{"one element, in filter", []int{3}, []int{}},
	{"several elements, all elements are distinct, filtered", []int{1, 2, 3}, []int{2}},
	{
		"several elements, duplicated elements, filtered",
		[]int{-5, -4, -3, -3, -2, -2, 0, 0, 1, 2, 3, 4, 5, 6, 6, 7, 7, 7},
		[]int{-4, -2, -2, 2, 4, 6, 6},
	},
	{"several elements, all not in list", []int{-6, -4, -2, 2, 4, 6, 8, 10}, []int{-6, -4, -2, 2, 4, 6, 8, 10}},
}

func TestSliceFilterByList(t *testing.T) {
	for _, testCase := range filterByListTestCases {
		result := SliceFilterByList(testCase.inputSlice)

		if len(result) != len(testCase.expected) {
			t.Fatalf(`%v. Expected length: %v. Actual length: %v`, testCase.caseName, len(testCase.expected), len(result))
		}

		for i, element := range result {
			if element != testCase.expected[i] {
				t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
			}
		}
	}
}
