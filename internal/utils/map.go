package utils

import "sort"

func MapInvert(inputMap map[string]int) map[int]string {
	result := make(map[int]string, 0)

	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		result[inputMap[key]] = key
	}

	return result
}
