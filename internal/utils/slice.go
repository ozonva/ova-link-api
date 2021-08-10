package utils

func SliceChunk(inputSlice []int, size int) [][]int {
	sliceLength := len(inputSlice)
	if sliceLength == 0 || size <= 0 {
		return make([][]int, 0, 0)
	}

	capacity := sliceLength / size
	if sliceLength%size != 0 {
		capacity += 1
	}

	result := make([][]int, 0, capacity)
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
	result := make([]int, 0, len(inputSlice))
	list := []int{-7, -5, -3, -1, 0, 1, 3, 5, 7}
	filter := getFilterMap(list)

	for _, value := range inputSlice {
		if _, ok := filter[value]; !ok {
			result = append(result, value)
			continue
		}
	}

	return result
}

func getFilterMap(slice []int) map[int]bool {
	filter := make(map[int]bool, len(slice))
	for _, v := range slice {
		filter[v] = true
	}

	return filter
}
