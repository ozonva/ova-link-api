package utils

import (
	"errors"
	"testing"

	"github.com/ozonva/ova-link-api/internal/link"
)

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

var link1 = *link.New(1, 1, "https://test1.com")
var link2 = *link.New(2, 1, "https://test2.com")
var link3 = *link.New(3, 1, "https://test3.com")
var link4 = *link.New(4, 2, "https://test4.com")
var link5 = *link.New(5, 2, "https://test5.com")
var link6 = *link.New(6, 2, "https://test6.com")
var link7 = *link.New(7, 3, "https://test7.com")
var link8 = *link.New(8, 3, "https://test8.com")
var link9 = *link.New(9, 3, "https://test9.com")

var chunkLinkTestCases = []struct {
	caseName   string
	inputSlice []link.Link
	inputSize  uint
	expected   [][]link.Link
}{
	{"empty slice", []link.Link{}, 3, [][]link.Link{}},
	{"one element, chunked equally", []link.Link{link1}, 1, [][]link.Link{{link1}}},
	{"several elements, size is more than length", []link.Link{link1, link2, link3}, 5, [][]link.Link{{link1, link2, link3}}},
	{"several elements, size = 0", []link.Link{link1, link2, link3}, 0, [][]link.Link{}},
	{"odd amount of elements, chunked equally", []link.Link{link1, link2, link3, link4, link5, link6, link7, link8, link9}, 3, [][]link.Link{{link1, link2, link3}, {link4, link5, link6}, {link7, link8, link9}}},
	{"odd amount of elements, chunked unequally", []link.Link{link1, link2, link3, link4, link5, link6, link7}, 2, [][]link.Link{{link1, link2}, {link3, link4}, {link5, link6}, {link7}}},
	{"even amount of elements, chunked equally", []link.Link{link1, link2, link3, link4, link5, link6}, 2, [][]link.Link{{link1, link2}, {link3, link4}, {link5, link6}}},
	{"even amount of elements, chunked unequally", []link.Link{link1, link2, link3, link4, link5, link6}, 4, [][]link.Link{{link1, link2, link3, link4}, {link5, link6}}},
}

func TestSliceChunkLink(t *testing.T) {
	for _, testCase := range chunkLinkTestCases {
		result := SliceChunkLink(testCase.inputSlice, testCase.inputSize)

		if len(result) != len(testCase.expected) {
			t.Fatalf(`%v. Expected length: %v. Actual length: %v`, testCase.caseName, len(testCase.expected), len(result))
		}

		for i, chunk := range result {
			for j, element := range chunk {
				if !element.Equals(&testCase.expected[i][j]) {
					t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
				}
			}
		}
	}
}

var linkDuplicate = *link.New(1, 1, "https://test1.com")
var sliceLinkToMapLinkTestCases = []struct {
	caseName   string
	inputSlice []link.Link
	result     map[uint64]link.Link
	error      error
}{
	{"empty slice", []link.Link{}, make(map[uint64]link.Link, 0), nil},
	{"slice without duplicates", []link.Link{link1, link2, link3}, map[uint64]link.Link{1: link1, 2: link2, 3: link3}, nil},
	{"slice with duplicates", []link.Link{link1, link2, link3, linkDuplicate}, nil, errors.New("duplicate link id")},
}

func TestSliceLinkToMapLink(t *testing.T) {
	for _, testCase := range sliceLinkToMapLinkTestCases {
		result, err := SliceLinkToMapLink(testCase.inputSlice)
		if err != nil && testCase.error == nil {
			t.Fatalf(`%v. Error expected: %v. Error actual: %v`, testCase.caseName, testCase.error, err)
		}

		if err != nil && err.Error() != testCase.error.Error() {
			t.Fatalf(`%v. Error expected: %v. Error actual: %v`, testCase.caseName, testCase.error, err)
		}

		if err == nil && testCase.error == nil {
			for i, value := range result {
				if _, ok := testCase.result[i]; !ok {
					t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.result, result)
				}

				v := testCase.result[i]
				if !value.Equals(&v) {
					t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.result, result)
				}
			}
		}
	}
}
