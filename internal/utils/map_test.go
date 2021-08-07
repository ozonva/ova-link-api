package utils

import "testing"

var invertTestCases = []struct {
	caseName string
	inputMap map[string]int
	expected map[int]string
}{
	{"empty map", map[string]int{}, map[int]string{}},
	{
		"distinct values",
		map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
		map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
	},
	{
		"duplicated values, only last key will be used",
		map[string]int{"a": 1, "b": 2, "c": 2, "d": 4, "e": 4},
		map[int]string{1: "a", 2: "c", 4: "e"},
	},
	{
		"distinct values unsorted keys",
		map[string]int{"c": 3, "d": 4, "b": 2, "e": 5, "a": 1},
		map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
	},
	{
		"duplicated values unsorted keys, only last key will be used",
		map[string]int{"c": 2, "d": 4, "a": 1, "e": 4, "b": 2},
		map[int]string{1: "a", 2: "c", 4: "e"},
	},
}

func TestMapInvert(t *testing.T) {
	for _, testCase := range invertTestCases {
		result := MapInvert(testCase.inputMap)

		if len(result) != len(testCase.expected) {
			t.Fatalf(`%v. Expected length: %v. Actual length: %v`, testCase.caseName, len(testCase.expected), len(result))
		}

		for key, value := range result {
			if value != testCase.expected[key] {
				t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
			}
		}
	}
}
