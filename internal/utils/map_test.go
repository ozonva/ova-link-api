package utils

import (
	"errors"
	"testing"
)

var invertTestCases = []struct {
	caseName string
	inputMap map[string]int
	expected map[int]string
	err      error
}{
	{"empty map", map[string]int{}, map[int]string{}, nil},
	{
		"distinct values, sorted keys",
		map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
		map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
		nil,
	},
	{
		"distinct values, unsorted keys",
		map[string]int{"c": 2, "d": 4, "a": 1, "e": 3, "b": 5},
		map[int]string{1: "a", 5: "b", 2: "c", 4: "d", 3: "e"},
		nil,
	},
	{
		"duplicated values, error",
		map[string]int{"a": 1, "b": 2, "c": 2, "d": 4, "e": 4},
		nil,
		errors.New("duplicate key exists"),
	},
}

func TestMapInvert(t *testing.T) {
	for _, testCase := range invertTestCases {
		result, err := MapInvert(testCase.inputMap)
		if err != nil {
			if testCase.err == nil {
				t.Fatalf(`%v. Expected err: %v. Actual: %v`, testCase.caseName, testCase.err, err)
			}
			if !errors.As(testCase.err, &err) {
				t.Fatalf(`%v. Expected err: %v. Actual: %v`, testCase.caseName, testCase.err, err)
			}
		}

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
