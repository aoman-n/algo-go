package quiz

func exclude[T comparable](list []T, target T) []T {
	result := make([]T, 0, len(list)-1)
	for _, v := range list {
		if target != v {
			result = append(result, v)
		}
	}
	return result
}

func SnakeStringV1(chars string) [][]string {
	pos := func(i int) int {
		return i + 1
	}
	neg := func(i int) int {
		return i - 1
	}

	operator := neg

	result := [][]string{{}, {}, {}}
	resultIndexes := []int{0, 1, 2}
	currentIndex := 1

	for _, char := range chars {
		result[currentIndex] = append(result[currentIndex], string(char))
		rest := exclude(resultIndexes, currentIndex)
		for _, v := range rest {
			result[v] = append(result[v], " ")
		}

		if currentIndex <= 0 {
			operator = pos
		}
		if currentIndex >= len(result)-1 {
			operator = neg
		}
		currentIndex = operator(currentIndex)
	}

	return result
}

func SnakeStringV2(chars string, depth int) [][]string {
	pos := func(i int) int {
		return i + 1
	}
	neg := func(i int) int {
		return i - 1
	}

	result := make([][]string, 0, depth)
	resultIndexes := make([]int, 0, depth)
	for i := range make([]struct{}, depth) {
		result = append(result, []string{})
		resultIndexes = append(resultIndexes, i)
	}
	i := depth / 2
	operator := neg

	for _, char := range chars {
		result[i] = append(result[i], string(char))
		rest := exclude(resultIndexes, i)
		for _, v := range rest {
			result[v] = append(result[v], " ")
		}

		if i <= 0 {
			operator = pos
		}
		if i >= depth-1 {
			operator = neg
		}
		i = operator(i)
	}

	return result
}
