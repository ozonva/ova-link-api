package utils

func SliceChunk(inputSlice []int, size int) [][]int {
	sliceLength := len(inputSlice)
	if sliceLength == 0 || size <= 0 {
		return make([][]int, 0, 0)
	}

	result := make([][]int, 0, sliceLength/size)
	for from := 0; from < sliceLength; {
		to := from + size

		if to > sliceLength {
			result = append(result, inputSlice[from:])
		} else {
			result = append(result, inputSlice[from:to])
		}

		from = to
	}

	return result
}

func SliceFilterByList(inputSlice []int) []int {
	hash := make(map[int]bool, 0)
	result := make([]int, 0, len(inputSlice))
	list := []int{-7, -5, -3, -1, 0, 1, 3, 5, 7}

	for _, value := range inputSlice {
		if _, ok := hash[value]; ok {
			result = append(result, value)
			continue
		}

		if binarySearch(value, list) == -1 {
			result = append(result, value)
			hash[value] = true
		}
	}

	return result
}

func binarySearch(needle int, haystack []int) int {
	to := len(haystack)
	mid := -1
	for from := 0; from < to; {
		mid = (from + to) / 2
		if needle == haystack[mid] {
			return mid
		}

		if needle < haystack[mid] {
			to = mid
		} else {
			from = mid + 1
		}
	}

	return -1
}
